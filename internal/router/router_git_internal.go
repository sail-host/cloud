package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
)

type GitInternalRouter struct{}

func (r *GitInternalRouter) InitRouter(Router *echo.Group) {
	gitInternalRouter := Router.Group("/git-internal")

	gitInternalApi := apiV1.ApiGroupApp.BaseApi
	{
		gitInternalRouter.GET("/list/:id", gitInternalApi.GitInternalList)
	}
}
