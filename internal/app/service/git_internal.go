package service

import (
	"errors"
	"fmt"
	"path"
	"sync"

	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/git"

	githubP "github.com/google/go-github/v65/github"
)

const PER_PAGE = 10

type GitInternalService struct {
}

type IGitInternalService interface {
	GetRepos(id uint, page int) (*dto.GitInternalRepoResponse, error)
	GetRepo(id uint) (*githubP.Repository, error)
	GetLastCommitInBranch(id uint, owner, repo, branch string) (*githubP.RepositoryCommit, error)
	CloneRepo(id uint, repo, branch, uuid string) error
	CreateDeployment(id uint, repo, uuid string, deployment model.Deployment) (int64, error)
	UpdateDeploymentStatus(id uint, repo, status, message string, deploymentID int64) error
	SetRepoWebhook(project *model.Project) error
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
		github := git.NewGithub(gitModel.Token, gitModel.Owner)
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
				Owner:         *repo.Owner.Login,
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

func (s *GitInternalService) GetRepo(id uint) (*githubP.Repository, error) {
	var gitManager *git.GitManager
	gitModel, err := gitRepo.GetGitByID(id)
	if err != nil {
		return nil, err
	}

	switch gitModel.Type {
	case "github":
		github := git.NewGithub(gitModel.Token, gitModel.Owner)
		gitManager = git.NewGitManager(github)
		repo, err := gitManager.GetRepo(gitModel.Owner, gitModel.Name)
		if err != nil {
			return nil, err
		}

		return repo, nil
	}

	return nil, nil
}

func (s *GitInternalService) GetLastCommitInBranch(id uint, owner, repo, branch string) (*githubP.RepositoryCommit, error) {
	var gitManager *git.GitManager
	gitModel, err := gitRepo.GetGitByID(id)
	if err != nil {
		return nil, err
	}

	switch gitModel.Type {
	case "github":
		github := git.NewGithub(gitModel.Token, gitModel.Owner)
		gitManager = git.NewGitManager(github)
		commit, err := gitManager.GetLastCommitInBranch(owner, repo, branch)
		if err != nil {
			return nil, err
		}
		return commit, nil
	}

	return nil, nil
}

func (s *GitInternalService) CloneRepo(id uint, repo, branch, uuid string) error {
	var gitManager *git.GitManager
	gitModel, err := gitRepo.GetGitByID(id)
	if err != nil {
		return err
	}

	deployDir := path.Join(global.CONF.System.DeployDir, uuid)

	switch gitModel.Type {
	case "github":
		github := git.NewGithub(gitModel.Token, gitModel.Owner)
		gitManager = git.NewGitManager(github)
		err = gitManager.CloneRepo(gitModel.Owner, repo, deployDir, branch, gitModel.Token, gitModel.Owner)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GitInternalService) CreateDeployment(id uint, repo, uuid string, deployment model.Deployment) (int64, error) {
	var gitManager *git.GitManager
	gitModel, err := gitRepo.GetGitByID(id)
	if err != nil {
		return 0, err
	}

	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		return 0, err
	}

	switch gitModel.Type {
	case "github":
		github := git.NewGithub(gitModel.Token, gitModel.Owner)
		gitManager = git.NewGitManager(github)
		gitDeploymentID, err := gitManager.CreateDeployment(gitModel.Owner, repo, &githubP.DeploymentRequest{
			Ref:         &project.ProductionBranch,
			Description: &deployment.GitMessage,
			Payload:     &deployment.GitAuthor,
		})
		if err != nil {
			return 0, err
		}

		return gitDeploymentID, nil
	}

	return 0, nil
}

func (s *GitInternalService) UpdateDeploymentStatus(id uint, repo, status, message string, deploymentID int64) error {
	var gitManager *git.GitManager
	gitModel, err := gitRepo.GetGitByID(id)
	if err != nil {
		return err
	}

	switch gitModel.Type {
	case "github":
		github := git.NewGithub(gitModel.Token, gitModel.Owner)
		gitManager = git.NewGitManager(github)
		err = gitManager.UpdateDeploymentStatus(gitModel.Owner, repo, status, message, deploymentID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GitInternalService) SetRepoWebhook(project *model.Project) error {
	webhookUrl := fmt.Sprintf("%s/api/v1/webhook/github/%d", global.CONF.System.AppUrl, project.ID)

	var gitManager *git.GitManager
	gitModel, err := gitRepo.GetGitByID(project.GitID)
	if err != nil {
		return err
	}

	switch gitModel.Type {
	case "github":
		github := git.NewGithub(gitModel.Token, gitModel.Owner)
		gitManager = git.NewGitManager(github)
		err = gitManager.SetRepoWebhook(gitModel.Owner, project.GitRepo, webhookUrl)
		if err != nil {
			return err
		}
	}

	return nil
}
