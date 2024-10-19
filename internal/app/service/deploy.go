package service

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/nodejs"
	"github.com/sail-host/cloud/internal/utils/sailhost"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/util/rand"
)

type DeployService struct{}

type IDeployService interface {
	CreateProject(c echo.Context, project *dto.CreateProjectRequest) error
	Deploy(project *model.Project)
	ListProjects() (*dto.ListProjectResponse, error)
	GetProjectWithName(projectName string) (*dto.BaseResponse, error)
	CheckProjectName(projectName string) (*dto.BaseResponse, error)
}

func NewIDeployService() IDeployService {
	return &DeployService{}
}

func (d *DeployService) CreateProject(c echo.Context, project *dto.CreateProjectRequest) error {
	err := c.Validate(project)
	if err != nil {
		return err
	}

	projectModel, err := projectRepo.CreateProject(&model.Project{
		Name:             project.Name,
		Framework:        project.Framework,
		GitUrl:           project.GitUrl,
		GitRepo:          project.GitRepo,
		ProductionBranch: project.ProductionBranch,
		GitID:            project.GitID,
		BuildCommand:     project.BuildCommand,
		OutputDir:        project.OutputDir,
		InstallCommand:   project.InstallCommand,
	})
	if err != nil {
		return err
	}

	global.LOG.Info("Project created and deploy started", projectModel)
	go d.Deploy(projectModel)

	return nil
}

func (d *DeployService) Deploy(project *model.Project) {
	global.LOG.Info("Deploying project", project)
	startTime := time.Now()

	// Get git info
	gitModel, err := gitRepo.GetGitByID(project.GitID)
	if err != nil {
		global.LOG.Error("Error getting git info", err)
		return
	}

	gitInternalService := NewIGitInternalService()

	// Get last commit
	commit, err := gitInternalService.GetLastCommitInBranch(gitModel.ID, gitModel.Owner, project.GitRepo, project.ProductionBranch)
	if err != nil {
		global.LOG.Error("Error getting last commit", err)
		return
	}

	// Create new deployment
	deployment := &model.Deployment{
		ProjectID:      project.ID,
		Status:         "pending",
		UUID:           uuid.New().String(),
		Active:         true,
		Ready:          false,
		DeploymentTime: 0,
		DeploymentSize: 0,

		// Git info
		GitHash:     commit.GetSHA(),
		GitMessage:  commit.GetCommit().GetMessage(),
		GitAuthor:   commit.GetCommit().GetAuthor().GetName(),
		GitCommitID: commit.GetCommit().GetTree().GetSHA(),
		GitDate:     commit.GetCommit().GetAuthor().GetDate().Time,
	}
	deployment, err = projectRepo.CreateDeployment(deployment)
	if err != nil {
		global.LOG.Error("Error creating deployment", err)
		return
	}

	// Create new log
	global.LOG.Info("Creating new log", deployment)
	err = projectRepo.CreateLog(deployment, "Clone git repository.")
	if err != nil {
		global.LOG.Error("Error creating deployment log", err)
		return
	}

	err = gitInternalService.CloneRepo(gitModel.ID, project.GitRepo, project.ProductionBranch, deployment.UUID)
	if err != nil {
		global.LOG.Error("Error cloning repo", err)
		return
	}

	err = projectRepo.CreateLog(deployment, "Git respository success clone.")
	if err != nil {
		global.LOG.Error("Error creating deployment log", err)
		return
	}

	// Check if node is installed
	nodeManager := nodejs.NewNodejsManager("v20", global.CONF.System.UtilsDir)
	exists, err := nodeManager.CheckVersionExist()
	if err != nil {
		global.LOG.Error("Error checking Nodejs version", err)
		return
	}
	if !exists {
		err = projectRepo.CreateLog(deployment, "Nodejs version not exists. Installing...")
		if err != nil {
			global.LOG.Error("Error creating deployment log", err)
			return
		}
		err = nodeManager.InstallVersion()
		if err != nil {
			global.LOG.Error("Error installing Nodejs", err)
			return
		}
	}
	runPath := path.Join(global.CONF.System.DeployDir, deployment.UUID)

	packageManager := nodejs.NewNodejsPackageManager(runPath)
	managers, err := packageManager.Check()
	if err != nil {
		global.LOG.Error("Error checking nodejs package manager", err)
		return
	}

	// Check user write install command
	if project.InstallCommand != "" {
		_, err = nodeManager.CmdNpmRun(project.InstallCommand, runPath)
		if err != nil {
			global.LOG.Error("Error running install command", err)
			return
		}
	} else {
		// Check nodejs package manager

		switch managers.Manager[0] {
		case "bun":
			_, err = nodeManager.CmdBunRun("install", runPath)
		case "pnpm":
			_, err = nodeManager.CmdPnpmRun("install", runPath)
		case "yarn":
			_, err = nodeManager.CmdYarnRun("install", runPath)
		default:
			_, err = nodeManager.CmdNpmRun("install", runPath)
		}
		if err != nil {
			global.LOG.Error("Error installing nodejs package manager", err)
			return
		}
	}

	if project.BuildCommand != "" {
		_, err = nodeManager.CmdNpmRun(project.BuildCommand, runPath)
		if err != nil {
			global.LOG.Error("Error running build command", err)
			return
		}
	} else {
		switch managers.Manager[0] {
		case "bun":
			_, err = nodeManager.CmdBunRun("run build", runPath)
		case "pnpm":
			_, err = nodeManager.CmdPnpmRun("run build", runPath)
		case "yarn":
			_, err = nodeManager.CmdYarnRun("run build", runPath)
		default:
			_, err = nodeManager.CmdNpmRun("run build", runPath)
		}
		if err != nil {
			global.LOG.Error("Error running build command", err)
			return
		}
	}

	// Generate domain
	domain := generateDomain(project.Name)

	// Create domain for project
	projectDomain := &model.ProjectDomain{
		ProjectID:  project.ID,
		Domain:     domain,
		Configured: true,
		Valid:      false,
	}
	_, err = projectRepo.CreateProjectDomain(projectDomain)
	if err != nil {
		global.LOG.Error("Error creating project domain", err)
		return
	}
	// Configure domain
	err = sailhost.ConfigureDomain(domain)
	if err != nil {
		global.LOG.Error("Error configuring domain", err)
		return
	}

	// Update deployment
	deployment.Status = "success"
	deployment.DeploymentTime = uint(time.Since(startTime).Seconds())
	// Calculate deployment size
	deploymentPath := path.Join(global.CONF.System.DeployDir, deployment.UUID, "dist")
	var size int64
	err = filepath.Walk(deploymentPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	if err != nil {
		global.LOG.Error("Error calculating deployment size", err)
	}
	deployment.DeploymentSize = uint(size)

	deployment.Ready = true
	err = projectRepo.UpdateDeployment(deployment)
	if err != nil {
		global.LOG.Error("Error updating deployment", err)
		return
	}

	// TODO: Write new service if exists run command

	// TODO: Write new nginx config for deployment

	// TODO: Restart nginx

	// TODO: Create new record for cloudflare dns if domain

	// TODO: Complete deployment

	global.LOG.Info("Deployment completed", deployment)
}

func generateDomain(projectName string) string {
	// Generate domain
	// Remove special characters and symbols from project name
	sanitizedName := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(projectName, "")
	// Ensure the domain name is not empty and starts with a letter or number
	if sanitizedName == "" || !regexp.MustCompile(`^[a-zA-Z0-9]`).MatchString(sanitizedName) {
		sanitizedName = "project" + sanitizedName
	}
	domain := fmt.Sprintf("%s.%s", sanitizedName, "sailhost.app")

	// Check if domain is already used
	used, err := sailhost.CheckDomainUsed(domain)
	if err != nil {
		global.LOG.Error("Error checking domain", err)
		panic(err)
	}
	// If domain is used, generate new domain
	if used {
		randomString := rand.String(5)
		domain = fmt.Sprintf("%s-%s.%s", sanitizedName, randomString, "sailhost.app")
	}

	return domain
}

func (d *DeployService) ListProjects() (*dto.ListProjectResponse, error) {
	projects, err := projectRepo.ListProjects()
	if err != nil {
		return nil, err
	}

	var projectListResponse []*dto.ProjectListResponse
	for _, project := range projects {
		lastDomain, err := projectRepo.GetLastDomain(project.ID)
		if err != nil {
			return nil, err
		}
		lastDeployment, err := projectRepo.GetLastDeployment(project.ID)
		if err != nil {
			return nil, err
		}
		projectListResponse = append(projectListResponse, &dto.ProjectListResponse{
			ID:        project.ID,
			Name:      project.Name,
			Domain:    lastDomain.Domain,
			GitHash:   lastDeployment.GitHash,
			GitDate:   lastDeployment.GitDate,
			GitBranch: project.ProductionBranch,
			GitCommit: lastDeployment.GitMessage,
		})
	}

	return &dto.ListProjectResponse{
		Status:  "success",
		Message: "Projects listed",
		Data:    projectListResponse,
	}, nil
}

func (d *DeployService) GetProjectWithName(projectName string) (*dto.BaseResponse, error) {
	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		return nil, err
	}

	var baseResponse dto.BaseResponse
	baseResponse.Status = "success"
	baseResponse.Message = "Project listed"
	baseResponse.Data = project

	return &baseResponse, nil
}

func (d *DeployService) CheckProjectName(projectName string) (*dto.BaseResponse, error) {
	var baseResponse dto.BaseResponse
	_, err := projectRepo.GetProjectWithName(projectName)
	baseResponse.Status = "success"

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			baseResponse.Message = "Project name is available"
			baseResponse.Data = true
			return &baseResponse, nil
		}
		baseResponse.Status = "error"
		baseResponse.Message = "Error checking project name"
		baseResponse.Data = false
		return &baseResponse, err
	}

	baseResponse.Message = "Project name is already used"
	baseResponse.Data = false

	return &baseResponse, nil
}
