package service

import (
	"context"

	"github.com/sail-host/cloud/internal/app/dto"
)

type ProjectSettingGeneralService struct{}

type IProjectSettingGeneralService interface {
	UpdateProjectName(ctx context.Context, projectName string, request dto.UpdateProjectNameRequest) error
	UpdateBuildAndOutputDir(ctx context.Context, projectName string, request dto.UpdateBuildAndOutputDirRequest) error
}

func NewIProjectSettingGeneralService() IProjectSettingGeneralService {
	return &ProjectSettingGeneralService{}
}

func (s *ProjectSettingGeneralService) UpdateProjectName(ctx context.Context, projectName string, request dto.UpdateProjectNameRequest) error {

	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		return err
	}

	return projectSettingRepo.UpdateProjectName(project.ID, request.Name)
}

func (s *ProjectSettingGeneralService) UpdateBuildAndOutputDir(ctx context.Context, projectName string, request dto.UpdateBuildAndOutputDirRequest) error {
	project, err := projectRepo.GetProjectWithName(projectName)
	if err != nil {
		return err
	}

	return projectSettingRepo.UpdateBuildAndOutputDir(project.ID, request.Framework, request.BuildCommand, request.OutputDir, request.InstallCommand)
}
