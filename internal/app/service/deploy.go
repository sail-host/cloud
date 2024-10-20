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
	})
	if err != nil {
		return err
	}

	global.LOG.Info("Project created and deploy started", projectModel)
	go d.Deploy(projectModel)

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
		return
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

	err = gitInternalService.CloneRepo(gitModel.ID, project.GitRepo, project.ProductionBranch, deployment.UUID)
	if err != nil {
		global.LOG.Error("Error cloning repo", err)
		return
	}

	err = projectRepo.CreateLog(deployment, "Git repository cloned.")
	if err != nil {
		global.LOG.Error("Error creating deployment log", err)
		return
	}

	nodejsDeploymentService := NewINodejsDeploymentService()

	err = nodejsDeploymentService.InstallDependencies(deployment)
	if err != nil {
		global.LOG.Error("Error installing dependencies", err)
		return
	}

	err = nodejsDeploymentService.Build(deployment)
	if err != nil {
		global.LOG.Error("Error building project", err)
		return
	}

	deploymentDomainService := NewIDeploymentDomainService()

	if !isRedeploy {
		err = deploymentDomainService.CreateSailHostDomain(deployment)
		if err != nil {
			global.LOG.Error("Error creating deployment domain", err)
			return
		}
	}

	// Update deployment
	deployment.Status = "success"
	deployment.DeploymentTime = uint(time.Since(startTime).Seconds())

	// TODO: Calculate deployment size
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

	// Update other deployments to not be current
	err = projectRepo.UpdateDeploymentIsCurrent(deployment.ID)
	if err != nil {
		global.LOG.Error("Error updating deployment is current", err)
		return
	}

	// TODO: Write new service if exists run command

	// Write new nginx config for deployment
	if !isRedeploy {
		nginxService := NewINginxService()
		err = nginxService.CreateNginxConfig(deployment)
		if err != nil {
			global.LOG.Error("Error creating nginx config", err)
			return
		}
	}

	// TODO: Restart nginx

	// TODO: Create new record for cloudflare dns if domain

	// TODO: Complete deployment

	global.LOG.Info("Deployment completed", deployment)
}
