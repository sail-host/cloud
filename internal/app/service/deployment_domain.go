package service

import (
	"fmt"
	"regexp"

	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/sailhost"
	"k8s.io/apimachinery/pkg/util/rand"
)

type DeploymentDomainService struct{}

type IDeploymentDomainService interface {
	CreateSailHostDomain(deployment *model.Deployment) error
	// TODO: Add other methods
}

func NewIDeploymentDomainService() IDeploymentDomainService {
	return &DeploymentDomainService{}
}

func (d *DeploymentDomainService) CreateSailHostDomain(deployment *model.Deployment) error {
	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		global.LOG.Error("Error getting project", err)
		return err
	}

	// Generate domain
	domain := generateDomain(project.Name)

	// Create domain for project
	projectDomain := &model.ProjectDomain{
		ProjectID:  project.ID,
		Project:    *project,
		Domain:     domain,
		Configured: true,
		Valid:      false,
	}
	_, err = projectRepo.CreateProjectDomain(projectDomain)
	if err != nil {
		global.LOG.Error("Error creating project domain", err)
		return err
	}
	// Configure domain
	err = sailhost.ConfigureDomain(domain)
	if err != nil {
		global.LOG.Error("Error configuring domain", err)
		return err
	}

	return nil
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
