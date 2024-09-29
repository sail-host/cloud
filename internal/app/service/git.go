package service

import (
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
)

type GitService struct {
}

type IGitService interface {
	GetGitByID(id uint) (*dto.GitResponse, *dto.BaseError)
	GetListGit() (*dto.GitListResponse, *dto.BaseError)
	CreateGit(c echo.Context, request dto.CreateGitRequest) (*dto.GitResponse, *dto.BaseError)
	UpdateGit(c echo.Context, id uint, request dto.UpdateGitRequest) (*dto.GitResponse, *dto.BaseError)
	DeleteGit(id uint) *dto.BaseError
}

func NewIGitService() IGitService {
	return &GitService{}
}

func (s *GitService) GetGitByID(id uint) (*dto.GitResponse, *dto.BaseError) {
	var response dto.GitResponse
	var baseError dto.BaseError

	git, err := gitRepo.GetGitByID(id)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Git not found"
		return nil, &baseError
	}

	response.Status = "success"
	response.Message = "Git retrieved successfully"
	response.Data.ID = git.ID
	response.Data.Name = git.Name
	response.Data.Url = git.Url
	response.Data.Type = git.Type
	response.Data.Token = git.Token

	return &response, nil
}

func (s *GitService) GetListGit() (*dto.GitListResponse, *dto.BaseError) {
	var response dto.GitListResponse
	var baseError dto.BaseError

	gitList, err := gitRepo.GetListGit()
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Git not found"
		return nil, &baseError
	}

	response.Status = "success"
	response.Message = "Git list retrieved successfully"
	response.Data = gitList

	return &response, nil
}

func (s *GitService) CreateGit(c echo.Context, request dto.CreateGitRequest) (*dto.GitResponse, *dto.BaseError) {
	var response dto.GitResponse
	var baseError dto.BaseError

	if err := c.Validate(request); err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return nil, &baseError
	}

	git := model.Git{
		Name:  request.Name,
		Url:   request.Url,
		Type:  request.Type,
		Token: request.Token,
	}

	err := gitRepo.CreateGit(&git)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Failed to create git"
		return nil, &baseError
	}

	response.Status = "success"
	response.Message = "Git created successfully"
	response.Data.ID = git.ID
	response.Data.Name = git.Name
	response.Data.Url = git.Url
	response.Data.Type = git.Type
	response.Data.Token = git.Token

	return &response, nil
}

func (s *GitService) UpdateGit(c echo.Context, id uint, request dto.UpdateGitRequest) (*dto.GitResponse, *dto.BaseError) {
	var response dto.GitResponse
	var baseError dto.BaseError

	git, err := gitRepo.GetGitByID(id)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Git not found"
		return nil, &baseError
	}

	if err := c.Validate(request); err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return nil, &baseError
	}

	git.Name = request.Name
	git.Url = request.Url
	git.Type = request.Type
	git.Token = request.Token

	err = gitRepo.UpdateGit(git)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Failed to update git"
		return nil, &baseError
	}

	response.Status = "success"
	response.Message = "Git updated successfully"
	response.Data.ID = git.ID
	response.Data.Name = git.Name
	response.Data.Url = git.Url
	response.Data.Type = git.Type
	response.Data.Token = git.Token

	return &response, nil
}

func (s *GitService) DeleteGit(id uint) *dto.BaseError {
	var baseError dto.BaseError

	git, err := gitRepo.GetGitByID(id)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Git not found"
		return &baseError
	}

	err = gitRepo.DeleteGit(git.ID)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Failed to delete git"
		return &baseError
	}

	return nil
}