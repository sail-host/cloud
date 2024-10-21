package service

import (
	"fmt"
	"path"
	"regexp"

	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/caddy"
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

	// Create project root directory
	publicPath := "dist"
	// TODO: Update this code using framework build path
	if project.Framework == "nextjs" {
		publicPath = ""
	} else if project.Framework == "react" {
		publicPath = "build"
	} else if project.Framework == "vue" {
		publicPath = "dist"
	} else if project.Framework == "svelte" {
		publicPath = "build"
	} else if project.Framework == "vite" {
		publicPath = "dist"
	}

	rootPath := path.Join(global.CONF.System.DeployDir, deployment.UUID, publicPath)

	// Create Caddy site
	// TODO: Update Caddy URL using global config
	webServer := caddy.NewCaddy("localhost:2019")
	err = webServer.CreateSite(&caddy.SiteConfig{
		Domain: domain,
		Root:   rootPath,
		SSL:    false,
	})
	if err != nil {
		global.LOG.Error("Error creating Caddy site", err)
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

	// Production domain
	sailhostDomain := "sailhost.app"
	if global.CONF.System.Mode == "dev" {
		sailhostDomain = "sailhost.local"
	}

	domain := fmt.Sprintf("%s.%s", sanitizedName, sailhostDomain)

	// Check if domain is already used
	used, err := sailhost.CheckDomainUsed(domain)
	if err != nil {
		global.LOG.Error("Error checking domain", err)
		panic(err)
	}
	// If domain is used, generate new domain
	if used {
		randomString := rand.String(5)
		domain = fmt.Sprintf("%s-%s.%s", sanitizedName, randomString, sailhostDomain)
	}

	return domain
}
