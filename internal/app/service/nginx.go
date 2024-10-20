package service

import (
	"os"

	"github.com/sail-host/cloud/internal/app/model"
)

type NginxService struct{}

type INginxService interface {
	CreateNginxConfig(deployment *model.Deployment) error
	UpdateNginxConfig(deployment *model.Deployment) error
}

func NewINginxService() INginxService {
	return &NginxService{}
}

func (n *NginxService) CreateNginxConfig(deployment *model.Deployment) error {
	// TODO: Create nginx config for deployment
	return nil
}

func (n *NginxService) UpdateNginxConfig(deployment *model.Deployment) error {
	return nil
}

func getConfigPath() string {
	// Check nginx config folder for common Linux distributions: Ubuntu, Debian, CentOS, etc.
	commonPaths := []string{
		"/etc/nginx/sites-available",
		"/etc/nginx/conf.d",
	}

	for _, path := range commonPaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// If no common path is found, return a default path
	return "/etc/nginx/sites-available"
}
