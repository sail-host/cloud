package dto

import "time"

type CreateProjectRequest struct {
	Name             string `json:"name" validate:"required"`
	Framework        string `json:"framework" validate:"required"`
	GitUrl           string `json:"git_url" validate:"required"`
	ProductionBranch string `json:"production_branch" validate:"required"`
	GitID            uint   `json:"git_id" validate:"required"`
	GitRepo          string `json:"git_repo" validate:"required"`
	BuildCommand     string `json:"build_command"`
	OutputDir        string `json:"output_dir"`
	InstallCommand   string `json:"install_command"`
	// TODO: Check request dto!
}

type ListProjectResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    []*ProjectListResponse `json:"data"`
}

type ProjectListResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	GitHash   string    `json:"git_hash"`
	GitDate   time.Time `json:"git_date"`
	GitBranch string    `json:"git_branch"`
	GitCommit string    `json:"git_commit"`
}

type GetProjectResponse struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	GitBranch string       `json:"git_branch"`
	GitCommit string       `json:"git_commit"`
	GitUrl    string       `json:"git_url"`
	GitHash   string       `json:"git_hash"`
	Domains   []DomainList `json:"domains"`
}

type DomainList struct {
	ID           uint      `json:"id"`
	Domain       string    `json:"domain"`
	IsDeployment bool      `json:"is_deployment"`
	CreatedAt    time.Time `json:"created_at"`
}

type ListDeploymentResponse struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	GitHash   string    `json:"git_hash"`
	GitCommit string    `json:"git_commit"`
	GitBranch string    `json:"git_branch"`
	GitDate   time.Time `json:"git_date"`
	GitUrl    string    `json:"git_url"`
	IsCurrent bool      `json:"is_current"`
	Size      int64     `json:"size"`
	User      string    `json:"user"`
}
