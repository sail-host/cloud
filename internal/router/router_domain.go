package router

import (
	"github.com/labstack/echo/v4"
	apiV1 "github.com/sail-host/cloud/internal/app/api/v1"
	"github.com/sail-host/cloud/internal/middleware"
)

type DomainRouter struct{}

func (r *DomainRouter) InitRouter(Router *echo.Group) {
	domainRouter := Router.Group("/domain")

	domainRouter.Use(middleware.AuthJWT)
	domainApi := apiV1.ApiGroupApp.BaseApi
	{
		domainRouter.GET("/list", domainApi.GetListDomain)
		domainRouter.POST("/create", domainApi.CreateDomain)
		domainRouter.PUT("/update/:id", domainApi.UpdateDomain)
		domainRouter.DELETE("/delete/:id", domainApi.DeleteDomain)
		domainRouter.GET("/show/:id", domainApi.GetDomainByID)
		domainRouter.POST("/check", domainApi.CheckDomain)
	}
}
