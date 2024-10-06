package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
)

type ProjectRouter struct{}

func (r *ProjectRouter) InitRouter(Router *echo.Group) {
	projectRouter := Router.Group("/project")

	projectApi := apiV1.ApiGroupApp.BaseApi
	{
		projectRouter.POST("/create", projectApi.CreateProject)
	}
}
