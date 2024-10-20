package model

import "time"

type Project struct {
	BaseModel
	Name      string `json:"name" gorm:"unique"`
	Framework string `json:"framework"`

	// Git info
	GitUrl           string `json:"git_url"`
	GitRepo          string `json:"git_repo"`
	ProductionBranch string `json:"production_branch"`
	GitID            uint   `json:"git_id"`
	Git              Git    `json:"git" gorm:"foreignKey:GitID"`

	// Build and output options
	BuildCommand   string `json:"build_command" gorm:"nullable"`
	OutputDir      string `json:"output_dir" gorm:"nullable"`
	InstallCommand string `json:"install_command" gorm:"nullable"`
}

// Environment variables
type EnvironmentVariable struct {
	BaseModel
	Name      string  `json:"name"`
	Value     string  `json:"value"`
	ProjectID uint    `json:"project_id"`
	Project   Project `json:"project" gorm:"foreignKey:ProjectID"`
}

// Deployment
type Deployment struct {
	BaseModel
	ProjectID      uint    `json:"project_id"`
	Project        Project `json:"project" gorm:"foreignKey:ProjectID"`
	UUID           string  `json:"uuid"`
	Status         string  `json:"status" gorm:"enum('pending', 'building', 'deploying', 'running', 'error', 'success')"`
	DeploymentTime uint    `json:"deployment_time"` // Time in seconds
	DeploymentSize uint    `json:"deployment_size"` // Size in bytes
	Active         bool    `json:"active"`
	Ready          bool    `json:"ready"`
	IsCurrent      bool    `json:"is_current"`

	// Git info
	GitHash     string    `json:"git_hash"`
	GitMessage  string    `json:"git_message"`
	GitAuthor   string    `json:"git_author"`
	GitCommitID string    `json:"git_commit_id"`
	GitDate     time.Time `json:"git_date"`
}

// Log
type Log struct {
	BaseModel
	DeploymentID uint       `json:"deployment_id"`
	Deployment   Deployment `json:"deployment" gorm:"foreignKey:DeploymentID"`
	Message      string     `json:"message"`
}

// Domain
type ProjectDomain struct {
	BaseModel
	ProjectID  uint    `json:"project_id"`
	Project    Project `json:"project" gorm:"foreignKey:ProjectID"`
	DomainID   uint    `json:"domain_id" gorm:"nullable"`
	Domain     string  `json:"domain"`
	Valid      bool    `json:"valid"`
	Configured bool    `json:"configured"`
}
