package service

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/sailhost"
	"k8s.io/apimachinery/pkg/util/rand"
)

type DeployService struct{}

type IDeployService interface {
	CreateProject(c echo.Context, project *dto.CreateProjectRequest) error
	Deploy(project *model.Project)
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
	commit, err := gitInternalService.GetLastCommitInBranch(gitModel.ID, project.ProductionBranch)
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

	// Clone repo
	// TODO: Clone repo

	// Run install command
	// TODO: Run install command

	// Run build command
	// TODO: Run build command

	// Generate domain
	domain := generateDomain(project.Name)

	// Create domain for project
	projectDomain := &model.ProjectDomain{
		ProjectID:  project.ID,
		Domain:     domain,
		Configured: true,
		Valid:      false,
	}
	projectDomain, err = projectRepo.CreateProjectDomain(projectDomain)
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

	// Complete
	// Update deployment
	deployment.Status = "success"
	deployment.DeploymentTime = uint(time.Since(startTime).Seconds())
	// TODO: Get deployment size
	deployment.DeploymentSize = 0
	deployment.Ready = true
	err = projectRepo.UpdateDeployment(deployment)
	if err != nil {
		global.LOG.Error("Error updating deployment", err)
		return
	}
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
