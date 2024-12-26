package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type SystemRouter struct {
}

func (r *SystemRouter) InitRouter(Router *echo.Group) {
	system := Router.Group("/system")

	system.Use(middleware.AuthJWT)
	systemApi := apiV1.ApiGroupApp.BaseApi
	{
		system.GET("/metrics", systemApi.GetSystemMetrics)
	}
}
