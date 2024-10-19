package service

import (
	"errors"

	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/global"
	"gorm.io/gorm"
)

type ProjectService struct{}

type IProjectService interface {
	GetProjectWithName(projectName string) (*dto.BaseResponse, error)
	CheckProjectName(projectName string) (*dto.BaseResponse, error)
	ListProjects() (*dto.ListProjectResponse, error)
}

func NewIProjectService() IProjectService {
	return &ProjectService{}
}

func (p *ProjectService) GetProjectWithName(projectName string) (*dto.BaseResponse, error) {
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
