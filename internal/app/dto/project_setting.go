package dto

type UpdateProjectNameRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateBuildAndOutputDirRequest struct {
	Framework      string `json:"framework" validate:"required"`
	BuildCommand   string `json:"build_command" validate:"nullable"`
	OutputDir      string `json:"output_dir" validate:"required"`
	InstallCommand string `json:"install_command" validate:"nullable"`
}

type AddNewDomainRequest struct {
	Domain   string `json:"domain" validate:"required"`
	DomainID uint   `json:"domain_id" validate:"required"`
}
