package web

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/handlers"
	"github.com/sail-host/cloud/web"
)

func NewServer() {
	assets, err := web.Assets()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.GET("/*", handlers.SPA(assets))

	// Api routes
	router := NewRoute(e)
	router.Run()

	log.Printf("Server started on port %s", global.CONF.System.Port)
	e.Logger.Fatal(e.Start(":" + global.CONF.System.Port))
}
