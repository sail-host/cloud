package service

import (
	"errors"

	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/utils/git"
	"github.com/sail-host/cloud/internal/utils/git/github"
)

type GitInternalService struct {
}

type IGitInternalService interface {
	GetRepos(id uint) ([]dto.GitInternalRepo, error)
}

func NewIGitInternalService() IGitInternalService {
	return &GitInternalService{}
}

func (s *GitInternalService) GetRepos(id uint) ([]dto.GitInternalRepo, error) {
	var repos []dto.GitInternalRepo
	var gitManager *git.GitManager

	gitModel, err := gitRepo.GetGitByID(id)
	if err != nil {
		return nil, err
	}

	switch gitModel.Type {
	case "github":
		github := github.NewGithub(gitModel.Token)
		gitManager = git.NewGitManager(github)
	case "gitlab":
		return nil, errors.New("gitlab not supported")
	case "bitbucket":
		return nil, errors.New("bitbucket not supported")
	case "gitea":
		return nil, errors.New("gitea not supported")
	default:
		return nil, errors.New("unknown git service type")
	}

	// Get Repositories
	reposFull, err := gitManager.GetRepos()
	if err != nil {
		return nil, err
	}

	// Process the repositories
	for _, repo := range reposFull {
		repos = append(repos, dto.GitInternalRepo{
			ID:            *repo.ID,
			Name:          *repo.Name,
			FullName:      *repo.FullName,
			Description:   repo.GetDescription(),
			URL:           *repo.HTMLURL,
			DefaultBranch: repo.GetDefaultBranch(),
			CloneURL:      repo.GetCloneURL(),
			Private:       repo.GetPrivate(),
			CreatedAt:     repo.GetCreatedAt().Time,
			UpdatedAt:     repo.GetUpdatedAt().Time,
		})
	}

	return repos, nil
}
