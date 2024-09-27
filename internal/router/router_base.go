package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
)

type BaseRouter struct {
}

func (r *BaseRouter) InitRouter(Router *echo.Group) {
	baseRouter := Router.Group("/auth")
	baseApi := apiV1.ApiGroupApp.BaseApi
	{
		baseRouter.POST("/login", baseApi.Login)
	}
}
