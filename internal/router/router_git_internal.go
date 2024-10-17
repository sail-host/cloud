package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type GitInternalRouter struct{}

func (r *GitInternalRouter) InitRouter(Router *echo.Group) {
	gitInternalRouter := Router.Group("/git-internal")

	gitInternalRouter.Use(middleware.AuthJWT)
	gitInternalApi := apiV1.ApiGroupApp.BaseApi
	{
		gitInternalRouter.GET("/list/:id", gitInternalApi.GitInternalList)
	}
}
