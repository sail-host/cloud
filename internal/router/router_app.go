package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/app/dto"
)

type AppRouter struct {
}

func (r *AppRouter) InitRouter(Router *echo.Group) {
	Router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.BaseResponse{Status: "success", Message: "pong"})
	})

	upgradeRouter := Router.Group("/upgrade")
	{
		upgradeRouter.GET("/check", apiV1.ApiGroupApp.BaseApi.CheckUpgrade)
		upgradeRouter.POST("/update", apiV1.ApiGroupApp.BaseApi.Update)
	}
}
