package repository

import (
	"fmt"

	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type ProjectRepo struct{}

type IProjectRepo interface {
	CreateProject(project *model.Project) (*model.Project, error)
	GetProjectByID(id uint) (*model.Project, error)
	UpdateProject(project *model.Project) error
	ListProjects() ([]*model.Project, error)
	DeleteProject(id uint) error
	GetProjectWithName(name string) (*model.Project, error)

	CreateDeployment(deployment *model.Deployment) (*model.Deployment, error)
	GetDeploymentByID(id uint) (*model.Deployment, error)
	UpdateDeployment(deployment *model.Deployment) error
	DeleteDeployment(id uint) error
	GetLastDeployment(id uint) (*model.Deployment, error)

	CreateProjectDomain(projectDomain *model.ProjectDomain) (*model.ProjectDomain, error)
	GetProjectDomainByID(id uint) (*model.ProjectDomain, error)
	UpdateProjectDomain(projectDomain *model.ProjectDomain) error
	DeleteProjectDomain(id uint) error
	GetLastDomain(id uint) (*model.ProjectDomain, error)

	CreateLog(dep *model.Deployment, log ...string) error
}

func NewIProjectRepo() IProjectRepo {
	return &ProjectRepo{}
}

func (p *ProjectRepo) CreateProject(project *model.Project) (*model.Project, error) {
	db := global.DB
	err := db.Create(project).Error
	if err != nil {
		return nil, err
	}

	db.Last(project)
	return project, nil
}

func (p *ProjectRepo) GetProjectByID(id uint) (*model.Project, error) {
	var project model.Project
	db := global.DB
	err := db.Where("id = ?", id).First(&project).Error
	return &project, err
}

func (p *ProjectRepo) UpdateProject(project *model.Project) error {
	db := global.DB
	return db.Model(project).Updates(project).Error
}

func (p *ProjectRepo) DeleteProject(id uint) error {
	db := global.DB
	return db.Where("id = ?", id).Delete(&model.Project{}).Error
}

func (p *ProjectRepo) CreateDeployment(deployment *model.Deployment) (*model.Deployment, error) {
	db := global.DB
	err := db.Create(deployment).Error
	if err != nil {
		return nil, err
	}

	db.Last(deployment)
	return deployment, nil
}

func (p *ProjectRepo) GetDeploymentByID(id uint) (*model.Deployment, error) {
	var deployment model.Deployment
	db := global.DB
	err := db.Where("id = ?", id).First(&deployment).Error
	return &deployment, err
}

func (p *ProjectRepo) UpdateDeployment(deployment *model.Deployment) error {
	db := global.DB
	return db.Model(deployment).Updates(deployment).Error
}

func (p *ProjectRepo) DeleteDeployment(id uint) error {
	db := global.DB
	return db.Where("id = ?", id).Delete(&model.Deployment{}).Error
}

func (p *ProjectRepo) CreateProjectDomain(projectDomain *model.ProjectDomain) (*model.ProjectDomain, error) {
	db := global.DB
	err := db.Create(projectDomain).Error
	if err != nil {
		return nil, err
	}

	db.Last(projectDomain)
	return projectDomain, nil
}

func (p *ProjectRepo) DeleteProjectDomain(id uint) error {
	db := global.DB
	return db.Where("id = ?", id).Delete(&model.ProjectDomain{}).Error
}

func (p *ProjectRepo) GetProjectDomainByID(id uint) (*model.ProjectDomain, error) {
	var projectDomain model.ProjectDomain
	db := global.DB
	err := db.Where("id = ?", id).First(&projectDomain).Error
	return &projectDomain, err
}

func (p *ProjectRepo) UpdateProjectDomain(projectDomain *model.ProjectDomain) error {
	db := global.DB
	return db.Model(projectDomain).Updates(projectDomain).Error
}

func (p *ProjectRepo) ListProjects() ([]*model.Project, error) {
	var projects []*model.Project
	db := global.DB
	err := db.Find(&projects).Error
	return projects, err
}

func (p *ProjectRepo) GetLastDomain(id uint) (*model.ProjectDomain, error) {
	var projectDomain model.ProjectDomain
	db := global.DB
	err := db.Where("project_id = ?", id).Order("created_at DESC").First(&projectDomain).Error
	return &projectDomain, err
}

func (p *ProjectRepo) GetLastDeployment(id uint) (*model.Deployment, error) {
	var deployment model.Deployment
	db := global.DB
	err := db.Where("project_id = ?", id).Order("created_at DESC").First(&deployment).Error
	return &deployment, err
}

func (p *ProjectRepo) GetProjectWithName(name string) (*model.Project, error) {
	var project model.Project
	db := global.DB
	err := db.Where("name = ?", name).First(&project).Error
	return &project, err
}

func (p *ProjectRepo) CreateLog(dep *model.Deployment, log ...string) error {
	var logModel model.Log
	db := global.DB

	var message string
	for _, i := range log {
		message = fmt.Sprintf("%s\n%s", message, i)
	}

	logModel.Message = message
	logModel.DeploymentID = dep.ID
	global.LOG.Info(message)

	err := db.Create(logModel).Error
	if err != nil {
		return err
	}

	return nil
}
