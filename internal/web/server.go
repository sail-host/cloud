package web

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sail-host/cloud/config"
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
	e.StaticFS("/*", assets)

	// Api routes
	router := NewRoute(e)
	router.Run()

	log.Printf("Server started on port %s", config.GetConfig().Port)
	e.Logger.Fatal(e.Start(":" + config.GetConfig().Port))
}
