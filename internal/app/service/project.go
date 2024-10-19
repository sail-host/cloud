package service

import (
	"errors"
	"strings"

	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/global"
	"gorm.io/gorm"
)

type ProjectService struct{}

type IProjectService interface {
	GetProjectWithName(projectName string) (*dto.BaseResponse, error)
	CheckProjectName(projectName string) (*dto.BaseResponse, error)
	ListProjects() (*dto.ListProjectResponse, error)
	GetProjectDeployments(projectName string) ([]*dto.ListDeploymentResponse, error)
}

func NewIProjectService() IProjectService {
	return &ProjectService{}
}

func (p *ProjectService) GetProjectWithName(projectName string) (*dto.BaseResponse, error) {
	var projectResponse dto.GetProjectResponse
	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		return nil, err
	}

	lastDeployment, err := projectRepo.GetLastDeployment(project.ID)
	if err != nil {
		return nil, err
	}

	projectResponse.ID = project.ID
	projectResponse.Name = project.Name
	projectResponse.Status = lastDeployment.Status
	projectResponse.CreatedAt = project.CreatedAt
	projectResponse.GitBranch = project.ProductionBranch
	projectResponse.GitCommit = lastDeployment.GitMessage
	projectResponse.GitUrl = project.GitUrl + "/" + project.GitRepo
	projectResponse.GitHash = lastDeployment.GitHash

	domains, err := projectRepo.ListProjectDomains(project.ID)
	if err != nil {
		return nil, err
	}
	for _, domain := range domains {
		isDeployment := false
		// Check domain end with .sailhost.app
		if strings.HasSuffix(domain.Domain, ".sailhost.app") {
			isDeployment = true
		}

		projectResponse.Domains = append(projectResponse.Domains, dto.DomainList{
			ID:           domain.ID,
			Domain:       domain.Domain,
			IsDeployment: isDeployment,
			CreatedAt:    domain.CreatedAt,
		})
	}

	var baseResponse dto.BaseResponse
	baseResponse.Status = "success"
	baseResponse.Message = "Project listed"
	baseResponse.Data = projectResponse

	return &baseResponse, nil
}

func (p *ProjectService) CheckProjectName(projectName string) (*dto.BaseResponse, error) {
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

func (p *ProjectService) ListProjects() (*dto.ListProjectResponse, error) {
	projects, err := projectRepo.ListProjects()
	if err != nil {
		global.LOG.Error("Error listing projects", err)
		return nil, err
	}

	var projectListResponse []*dto.ProjectListResponse
	for _, project := range projects {
		lastDomain, err := projectRepo.GetLastDomain(project.ID)
		if err != nil {
			global.LOG.Error("Error getting last domain", err)
			return nil, err
		}
		lastDeployment, err := projectRepo.GetLastDeployment(project.ID)
		if err != nil {
			global.LOG.Error("Error getting last deployment", err)
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

func (p *ProjectService) GetProjectDeployments(projectName string) ([]*dto.ListDeploymentResponse, error) {
	var deploymentListResponse []*dto.ListDeploymentResponse
	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		return nil, err
	}
	deployments, err := projectRepo.ListProjectDeployments(project.ID)
	if err != nil {
		return nil, err
	}
	lastDeployment, err := projectRepo.GetLastDeployment(project.ID)
	if err != nil {
		return nil, err
	}

	for _, deployment := range deployments {
		deploymentListResponse = append(deploymentListResponse, &dto.ListDeploymentResponse{
			ID:        deployment.ID,
			Status:    deployment.Status,
			CreatedAt: deployment.CreatedAt,
			GitHash:   deployment.GitHash,
			GitCommit: deployment.GitMessage,
			GitBranch: project.ProductionBranch,
			GitDate:   deployment.GitDate,
			IsCurrent: deployment.ID == lastDeployment.ID,
			Size:      int64(deployment.DeploymentSize),
			User:      deployment.GitAuthor,
		})
	}

	return deploymentListResponse, nil
}
