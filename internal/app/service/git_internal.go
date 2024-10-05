package service

import (
	"errors"
	"sync"

	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/utils/git"
	"github.com/sail-host/cloud/internal/utils/git/github"

	githubP "github.com/google/go-github/v65/github"
)

const PER_PAGE = 10

type GitInternalService struct {
}

type IGitInternalService interface {
	GetRepos(id uint, page int) (*dto.GitInternalRepoResponse, error)
}

func NewIGitInternalService() IGitInternalService {
	return &GitInternalService{}
}

func (s *GitInternalService) GetRepos(id uint, page int) (*dto.GitInternalRepoResponse, error) {
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
	res, err := gitManager.GetRepos(page, PER_PAGE)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	repos := make([]dto.GitInternalRepo, 0, len(res.Repos))

	// Process the repositories
	for _, repo := range res.Repos {
		wg.Add(1)
		go func(repo *githubP.Repository) {
			defer wg.Done()
			var framework string
			framework, _ = gitManager.GetFramework(*repo.Owner.Login, *repo.Name)

			newRepo := dto.GitInternalRepo{
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
				Framework:     framework,
			}

			mu.Lock()
			repos = append(repos, newRepo)
			mu.Unlock()
		}(repo)
	}

	wg.Wait()

	var response dto.GitInternalRepoResponse
	response.Data = repos
	response.NextPage = res.NextPage
	response.LastPage = res.LastPage

	return &response, nil
}
