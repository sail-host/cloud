package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type ProjectRouter struct{}

func (r *ProjectRouter) InitRouter(Router *echo.Group) {
	projectRouter := Router.Group("/project")

	projectRouter.Use(middleware.AuthJWT)
	projectApi := apiV1.ApiGroupApp.BaseApi
	{
		projectRouter.POST("/create", projectApi.CreateProject)
		projectRouter.GET("/list", projectApi.ListProjects)
		projectRouter.GET("/show/:name", projectApi.GetProjectWithName)
		projectRouter.GET("/check", projectApi.CheckProjectName)
		projectRouter.GET("/deployments/:name", projectApi.GetProjectDeployments)
		projectRouter.POST("/redeploy/:name", projectApi.RedeployProject)
	}

	projectSettingRouter := Router.Group("/project-setting")
	projectSettingApi := apiV1.ApiGroupApp.BaseApi
	{
		projectSettingRouter.PUT("/update-name/:name", projectSettingApi.UpdateProjectName)
		projectSettingRouter.PUT("/update-build-and-output-dir/:name", projectSettingApi.UpdateBuildAndOutputDir)
	}
}
