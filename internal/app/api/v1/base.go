package v1

import "github.com/sail-host/cloud/internal/app/service"

type BaseApi struct{}

type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var (
	authService = service.NewIAuthService()
	gitService  = service.NewIGitService()
)
