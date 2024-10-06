package repository

import (
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type ProjectRepo struct {
}

type IProjectRepo interface {
	CreateProject(project *model.Project) error
	GetProjectByID(id uint) (*model.Project, error)
	UpdateProject(project *model.Project) error
	DeleteProject(id uint) error
}

func NewIProjectRepo() IProjectRepo {
	return &ProjectRepo{}
}

func (p *ProjectRepo) CreateProject(project *model.Project) error {
	db := global.DB
	return db.Create(project).Error
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
