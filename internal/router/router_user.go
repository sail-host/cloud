package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type UserRouter struct {
}

func (r *UserRouter) InitRouter(Router *echo.Group) {
	userRouter := Router.Group("/user")

	userRouter.Use(middleware.AuthJWT)
	userApi := apiV1.ApiGroupApp.BaseApi
	{
		userRouter.GET("", userApi.GetUser)
	}
}
