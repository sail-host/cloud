package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type BaseRouter struct {
}

func (r *BaseRouter) InitRouter(Router *echo.Group) {
	baseRouter := Router.Group("/auth")
	baseApi := apiV1.ApiGroupApp.BaseApi
	{
		baseRouter.POST("/login", baseApi.Login)
		baseRouter.GET("/check-user-first-time", baseApi.CheckUserFirstTime)
		baseRouter.POST("/register", baseApi.Register)
	}

	baseRouter.Use(middleware.AuthJWT)
	{
		baseRouter.POST("/logout", baseApi.Logout)
	}
}
