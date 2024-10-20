package service

import "github.com/sail-host/cloud/internal/app/model"

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
