package server

import (
	"github.com/sail-host/cloud/internal/init/app"
	"github.com/sail-host/cloud/internal/init/cache"
	"github.com/sail-host/cloud/internal/init/db"
	"github.com/sail-host/cloud/internal/init/log"
	"github.com/sail-host/cloud/internal/init/migration"
	"github.com/sail-host/cloud/internal/init/session"
	"github.com/sail-host/cloud/internal/init/validator"
	"github.com/sail-host/cloud/internal/init/viper"
)

func Init() {
	viper.Init()
	log.Init()
	db.Init()
	migration.Init()
	validator.Init()
	app.Init()
	cache.Init()
	session.Init()
}
