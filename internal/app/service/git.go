package service

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/utils/git"
	"github.com/sail-host/cloud/internal/utils/git/github"
)

type GitService struct {
}

type IGitService interface {
	GetGitByID(id uint) (*dto.GitResponse, *dto.BaseError)
	GetListGit() (*dto.GitListResponse, *dto.BaseError)
	CreateGit(c echo.Context, request dto.CreateGitRequest) (*dto.GitResponse, *dto.BaseError)
	UpdateGit(c echo.Context, id uint, request dto.UpdateGitRequest) (*dto.GitResponse, *dto.BaseError)
	DeleteGit(id uint) *dto.BaseError
	CheckAccount(c echo.Context, request dto.CreateGitRequest) (*dto.BaseResponse, *dto.BaseError)
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

	// Parse owner from url. Last /
	splitText := strings.Split(request.Url, "/")
	owner := splitText[len(splitText)-1]

	if owner == "" {
		baseError.Status = "error"
		baseError.Message = "Error in parsing owner from url"
		return nil, &baseError
	}

	git := model.Git{
		Name:  request.Name,
		Url:   request.Url,
		Type:  request.Type,
		Token: request.Token,
		Owner: owner,
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

	splitText := strings.Split(request.Url, "/")
	owner := splitText[len(splitText)-1]

	if owner == "" {
		baseError.Status = "error"
		baseError.Message = "Error in parsing owner from url"
		return nil, &baseError
	}

	git.Name = request.Name
	git.Url = request.Url
	git.Owner = owner
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

func (s *GitService) CheckAccount(c echo.Context, request dto.CreateGitRequest) (*dto.BaseResponse, *dto.BaseError) {
	var response dto.BaseResponse
	var baseError dto.BaseError
	response.Data = false

	if request.Type == "github" {
		github := github.NewGithub(request.Token)
		gitManager := git.NewGitManager(github)

		account, err := gitManager.CheckAccount()
		if err != nil {
			baseError.Status = "error"
			baseError.Message = "Failed to check account"
			return nil, &baseError
		}
		response.Data = account
	}
	// TODO: Add other git provider

	response.Status = "success"
	response.Message = "Git account checked successfully"

	return &response, nil
}
