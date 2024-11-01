package router

import (
	"time"

	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type BaseRouter struct {
}

func (r *BaseRouter) InitRouter(Router *echo.Group) {
	baseRouter := Router.Group("/auth")

	baseRouter.Use(middleware.RateLimiter(5, 30*time.Second)) // 5 requests per 30 seconds

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
