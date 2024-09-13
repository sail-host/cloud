package main

import (
	"github.com/sail-host/cloud/config"
	"github.com/sail-host/cloud/internal/bootstrap"
)

func main() {
	// Load config
	config.LoadConfig()

	// Run web server
	app := bootstrap.NewApp()
	app.Run()
}
