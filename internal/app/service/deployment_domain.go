package service

import (
	"fmt"
	"path"
	"regexp"

	sdn_cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/caddy"
	"github.com/sail-host/cloud/internal/utils/cloudflare"
	"github.com/sail-host/cloud/internal/utils/framework"
	"github.com/sail-host/cloud/internal/utils/ip"
	"github.com/sail-host/cloud/internal/utils/sailhost"
	"k8s.io/apimachinery/pkg/util/rand"
)

type DeploymentDomainService struct{}

type IDeploymentDomainService interface {
	CreateSailHostDomain(deployment *model.Deployment) error
	AddNewDomain(projectName string, domain dto.AddNewDomainRequest) (any, error)
	RemoveDomain(projectDomainID uint) error
	DomainsList(projectName string) (dto.BaseResponse, error)
	UpdateWebServerConfig(project *model.Project, deployment *model.Deployment) error
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
	var publicPath string
	if project.OutputDir != "" {
		publicPath = project.OutputDir
	} else {
		publicPath = framework.OutputDir(project.Framework)
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

func (d *DeploymentDomainService) AddNewDomain(projectName string, domain dto.AddNewDomainRequest) (any, error) {
	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		global.LOG.Error("Error getting project", err)
		return nil, err
	}

	// TODO: Update active and last deployment
	deployment, err := projectRepo.GetLastDeployment(project.ID)
	if err != nil {
		global.LOG.Error("Error getting deployment", err)
		return nil, err
	}

	domainModel, err := domainRepo.GetDomainByID(domain.DomainID)
	if err != nil {
		global.LOG.Error("Error getting domain", err)
		return nil, err
	}

	// Create domain
	var fullDomain string
	if domain.Domain == "@" {
		fullDomain = domainModel.Domain
	} else {
		fullDomain = fmt.Sprintf("%s.%s", domain.Domain, domainModel.Domain)
	}

	projectDomain := &model.ProjectDomain{
		ProjectID:  project.ID,
		Project:    *project,
		DomainID:   domainModel.ID,
		Domain:     fullDomain,
		Configured: false,
		Valid:      false,
	}
	_, err = projectRepo.CreateProjectDomain(projectDomain)
	if err != nil {
		global.LOG.Error("Error creating project domain", err)
		return nil, err
	}

	publicIP, err := ip.GetPublicIP()
	if err != nil {
		global.LOG.Error("Error getting public IP", err)
		return nil, err
	}

	// Check domain managed using cloudflare create new record
	if domainModel.DNSProvider == "cloudflare" {
		cloudflareManager, err := cloudflare.NewManager(domainModel.CloudflareAPIKey)
		if err != nil {
			global.LOG.Error("Error creating cloudflare manager", err)
			return nil, err
		}

		_, err = cloudflareManager.CreateDNSRecord(domainModel.CloudflareZoneID, sdn_cloudflare.CreateDNSRecordParams{
			Name:      fullDomain,
			Type:      "A",
			Content:   publicIP,
			Proxiable: true,
		})
		if err != nil {
			global.LOG.Error("Error creating cloudflare DNS record", err)
			return nil, err
		}
	}

	// Update project domain
	projectDomain.Configured = true
	projectDomain.Valid = true
	err = projectRepo.UpdateProjectDomain(projectDomain)
	if err != nil {
		global.LOG.Error("Error updating project domain", err)
		return nil, err
	}

	// Update web server config
	var publicPath string
	if project.OutputDir != "" {
		publicPath = project.OutputDir
	} else {
		publicPath = framework.OutputDir(project.Framework)
	}

	rootPath := path.Join(global.CONF.System.DeployDir, deployment.UUID, publicPath)

	// TODO: Update Caddy URL using global config
	webServer := caddy.NewCaddy("localhost:2019")
	err = webServer.CreateSite(&caddy.SiteConfig{
		Domain: fullDomain,
		Root:   rootPath,
		SSL:    domainModel.DNSProvider != "cloudflare",
	})
	if err != nil {
		global.LOG.Error("Error creating Caddy site", err)
		return nil, err
	}

	if domainModel.DNSProvider != "cloudflare" {
		resp := dto.AddNewDomainResponse{
			IP:         publicIP,
			FullDomain: fullDomain,
			Domain:     domain.Domain,
			Type:       "A",
		}
		return resp, nil
	}

	return nil, nil
}

func (d *DeploymentDomainService) RemoveDomain(projectDomainID uint) error {
	projectDomain, err := projectRepo.GetProjectDomainByID(projectDomainID)
	if err != nil {
		return err
	}

	// Remove Caddy site
	webServer := caddy.NewCaddy("localhost:2019")
	err = webServer.RemoveSite(projectDomain.Domain)
	if err != nil {
		return err
	}

	// Remove domain from project
	err = projectRepo.DeleteProjectDomain(projectDomainID)
	if err != nil {
		return err
	}

	return nil
}

func (d *DeploymentDomainService) DomainsList(projectName string) (dto.BaseResponse, error) {
	var resp dto.BaseResponse
	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		global.LOG.Error("Error getting project", err)
		resp.Message = err.Error()
		resp.Status = "error"
		return resp, err
	}

	domains, err := projectRepo.ListProjectDomains(project.ID)
	if err != nil {
		global.LOG.Error("Error getting project domains", err)
		resp.Message = err.Error()
		resp.Status = "error"
		return resp, err
	}

	resp.Data = domains
	resp.Status = "success"
	return resp, nil
}

func (d *DeploymentDomainService) UpdateWebServerConfig(project *model.Project, deployment *model.Deployment) error {

	// Get project domains
	domains, err := projectRepo.ListProjectDomains(project.ID)
	if err != nil {
		return err
	}

	// Update Caddy sites
	// TODO: Update Caddy URL using global config
	webServer := caddy.NewCaddy("localhost:2019")
	for _, domain := range domains {
		err = webServer.UpdateSite(&caddy.SiteConfig{
			Domain: domain.Domain,
			Root:   path.Join(global.CONF.System.DeployDir, deployment.UUID, domain.Project.OutputDir),
		})
		if err != nil {
			global.LOG.Error("Error updating Caddy site", err)
			continue
		}
	}

	return nil
}
