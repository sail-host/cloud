package bootstrap

import "github.com/sail-host/cloud/internal/web"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	// Run web server
	web.NewServer()
}
