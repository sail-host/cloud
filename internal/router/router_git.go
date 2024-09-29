package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type GitRouter struct{}

func (r *GitRouter) InitRouter(Router *echo.Group) {
	gitRouter := Router.Group("/git")

	gitRouter.Use(middleware.AuthJWT)
	gitApi := apiV1.ApiGroupApp.BaseApi
	{
		gitRouter.GET("/list", gitApi.GitList)
		gitRouter.POST("/create", gitApi.GitCreate)
		gitRouter.PUT("/update/:id", gitApi.GitUpdate)
		gitRouter.DELETE("/delete/:id", gitApi.GitDelete)
		gitRouter.GET("/show/:id", gitApi.GitGetByID)
		gitRouter.POST("/check-account", gitApi.GitCheckAccount)
	}
}
