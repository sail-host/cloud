package service

import (
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/framework"
	"github.com/sail-host/cloud/internal/utils/nodejs"
)

type DeployService struct{}

type IDeployService interface {
	CreateProject(c echo.Context, project *dto.CreateProjectRequest) error
	Deploy(project *model.Project)
	Redeploy(c echo.Context, projectName string) error
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
		RootDir:          project.RootDir,
		NodeVersion:      "v20", // Default node version
	})
	if err != nil {
		return err
	}

	global.LOG.Info("Project created and deploy started", projectModel)
	go d.Deploy(projectModel)

	// Set repo webhook
	gitInternalService := NewIGitInternalService()
	err = gitInternalService.SetRepoWebhook(projectModel)
	if err != nil {
		global.LOG.Error("Error setting repo webhook", err)
	}

	return nil
}

func (d *DeployService) Redeploy(c echo.Context, projectName string) error {
	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		return err
	}

	global.LOG.Info("Project redeploy started", project)
	go d.Deploy(project)

	return nil
}

func (d *DeployService) Deploy(project *model.Project) {
	global.LOG.Info("Deploying project", project)
	startTime := time.Now()
	lastDeployment, err := projectRepo.GetLastDeployment(project.ID)
	if err != nil {
		global.LOG.Error("Error getting last deployment", err)
	}

	isRedeploy := lastDeployment.ID != 0

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
		IsCurrent:      true,
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

	// Create new github deployment
	var gitDeploymentID int64
	gitDeploymentID, err = gitInternalService.CreateDeployment(gitModel.ID, project.GitRepo, deployment.UUID, *deployment)
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error creating github deployment", 0)
		return
	}

	gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "in_progress", "Deployment in progress", gitDeploymentID)

	// Clone repo
	err = gitInternalService.CloneRepo(gitModel.ID, project.GitRepo, project.ProductionBranch, deployment.UUID)
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error cloning repo", gitDeploymentID)
		return
	}

	err = projectRepo.CreateLog(deployment, "Git repository cloned.")
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error creating deployment log", gitDeploymentID)
		return
	}

	// Check node version
	nodeVersion, err := nodejs.GetVersion(path.Join(global.CONF.System.DeployDir, deployment.UUID))
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error getting node version", gitDeploymentID)
		return
	}
	if nodeVersion != "" {
		project.NodeVersion = nodeVersion
		err = projectRepo.UpdateProject(project)
		if err != nil {
			errorDeployment(deployment, err)
			gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error updating project node version", gitDeploymentID)
			return
		}
	}

	nodejsDeploymentService := NewINodejsDeploymentService()

	// Install dependencies
	err = nodejsDeploymentService.InstallDependencies(deployment)
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error installing dependencies", gitDeploymentID)
		return
	}

	// Build project
	err = nodejsDeploymentService.Build(deployment)
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error building project", gitDeploymentID)
		return
	}

	// Start project
	err = nodejsDeploymentService.Start(deployment)
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error starting project", gitDeploymentID)
		return
	}

	deploymentDomainService := NewIDeploymentDomainService()

	// Create sailhost domain
	if !isRedeploy {
		err = deploymentDomainService.CreateSailHostDomain(deployment)
		if err != nil {
			errorDeployment(deployment, err)
			gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error creating deployment domain", gitDeploymentID)
			return
		}
	}

	// Redeploy update web server config
	if isRedeploy {
		err = deploymentDomainService.UpdateWebServerConfig(project, deployment)
		if err != nil {
			errorDeployment(deployment, err)
			gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error updating web server config", gitDeploymentID)
			return
		}
	}

	// Update deployment
	deployment.Status = "success"
	deployment.DeploymentTime = uint(time.Since(startTime).Seconds())

	// Calculate deployment size
	var buildDir string
	if project.OutputDir != "" {
		buildDir = project.OutputDir
	} else {
		buildDir = framework.OutputDir(project.Framework)
	}
	deploymentPath := path.Join(global.CONF.System.DeployDir, deployment.UUID, project.RootDir, buildDir)
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
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error updating deployment", gitDeploymentID)
		return
	}

	// Update other deployments to not be current
	err = projectRepo.UpdateDeploymentIsCurrent(deployment.ID)
	if err != nil {
		errorDeployment(deployment, err)
		gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error updating deployment is current", gitDeploymentID)
		return
	}

	// Write new service if exists run command
	systemdService := NewIDeploymentSystemdService()
	if !isRedeploy {

		err = systemdService.CreateConfig(deployment, deploymentPath)
		if err != nil {
			errorDeployment(deployment, err)
			gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error creating systemd service", gitDeploymentID)
			return
		}
	} else {
		err = systemdService.RestartService(deployment, deploymentPath)
		if err != nil {
			errorDeployment(deployment, err)
			gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "failure", "Error restarting systemd service", gitDeploymentID)
			return
		}
	}

	global.LOG.Info("Deployment completed", deployment)
	gitInternalService.UpdateDeploymentStatus(gitModel.ID, project.GitRepo, "success", "Deployment completed", gitDeploymentID)
}

func errorDeployment(deployment *model.Deployment, err error) {
	global.LOG.Error("Error deploying project", err)

	deployment.Status = "error"
	deployment.Ready = false

	err = projectRepo.CreateLog(deployment, "Error deploying project", err.Error())
	if err != nil {
		global.LOG.Error("Error creating deployment log", err)
	}

	projectRepo.UpdateDeployment(deployment)
}
