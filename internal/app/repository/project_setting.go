package repository

import (
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type ProjectSettingRepository struct{}

type IProjectSettingRepository interface {
	UpdateProjectName(projectID uint, name string) error
	UpdateBuildAndOutputDir(projectID uint, framework string, buildCommand string, outputDir string, installCommand string) error
}

func NewIProjectSettingRepository() IProjectSettingRepository {
	return &ProjectSettingRepository{}
}

func (r *ProjectSettingRepository) UpdateProjectName(projectID uint, name string) error {
	db := global.DB
	return db.Model(&model.Project{}).Where("id = ?", projectID).Update("name", name).Error
}

func (r *ProjectSettingRepository) UpdateBuildAndOutputDir(projectID uint, framework string, buildCommand string, outputDir string, installCommand string) error {
	db := global.DB
	return db.Model(&model.Project{}).Where("id = ?", projectID).Updates(map[string]interface{}{
		"framework":       framework,
		"build_command":   buildCommand,
		"output_dir":      outputDir,
		"install_command": installCommand,
	}).Error
}
