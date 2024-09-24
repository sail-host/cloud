package main

import (
	"github.com/sail-host/cloud/internal/bootstrap"
)

func main() {
	// Run web server
	app := bootstrap.NewApp()
	app.Run()
}
