package router

import (
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sail-host/cloud/internal/app/service"
	"github.com/sail-host/cloud/internal/middleware"
	commRoutes "github.com/sail-host/cloud/internal/router"
	"github.com/sail-host/cloud/web"
)

var Router *echo.Echo

func setStatic(rootRouter *echo.Echo) {
	rootRouter.Static("/uploads", "./uploads")

	// Serve static files
	assets, err := web.Assets()
	if err != nil {
		log.Fatal(err)
	}
	rootRouter.GET("/*", service.NewSPA(assets).ServeHTTP())
}

func Routers() *echo.Echo {
	Router = echo.New()

	// TODO: Add this middleware for log and debug
	Router.Use(echoMiddleware.Logger())
	Router.Use(echoMiddleware.Recover())

	setStatic(Router)

	routeGroup := Router.Group("/api/v1")
	routeGroup.Use(middleware.ContentJSON)

	// Register common routes
	for _, route := range commRoutes.RouterGroupApp {
		route.InitRouter(routeGroup)
	}

	return Router
}
