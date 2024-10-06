package dto

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
