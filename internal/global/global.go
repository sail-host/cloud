package global

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/go-playground/validator/v10"
	"github.com/sail-host/cloud/internal/config"
	"github.com/sail-host/cloud/internal/init/cache/badger_db"
	"github.com/sail-host/cloud/internal/init/session/psession"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	MonitorDB *gorm.DB
	LOG       *logrus.Logger
	CONF      config.Config
	VALID     *validator.Validate
	SESSION   *psession.PSession
	CACHE     *badger_db.Cache
	CacheDb   *badger.DB
	Viper     *viper.Viper
)
