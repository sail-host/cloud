package viper

import (
	"github.com/sail-host/cloud/internal/config"
	"github.com/sail-host/cloud/internal/global"
	"github.com/spf13/viper"
)

func Init() {
	baseDir := "/opt"
	port := "8080"
	mode := "dev"
	version := "v0.0.1"
	v := viper.NewWithOptions()
	v.SetConfigType("yaml")

	// TODO: this test not work in production!
	serverConfig := config.Config{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	serverConfig.Debug = true
	serverConfig.Log.LogName = "cloud"
	serverConfig.Log.LogSuffix = ".log"
	serverConfig.Log.MaxBackup = 10
	serverConfig.Log.TimeZone = "Europe/Istanbul"
	serverConfig.Log.Level = "debug"

	serverConfig.System.BaseDir = baseDir
	serverConfig.System.Port = port
	serverConfig.System.Mode = mode
	serverConfig.System.Version = version
	serverConfig.System.LogPath = "/Users/whoami/Documents/files/golang/sailhost.io/cloud/tmp/logs"
	serverConfig.System.DbFile = "cloud.db"
	serverConfig.System.DbPath = "/Users/whoami/Documents/files/golang/sailhost.io/cloud/tmp/data"
	serverConfig.System.DataDir = "/Users/whoami/Documents/files/golang/sailhost.io/cloud/tmp/data"
	serverConfig.System.TmpDir = "/Users/whoami/Documents/files/golang/sailhost.io/cloud/tmp"
	serverConfig.System.Cache = "/Users/whoami/Documents/files/golang/sailhost.io/cloud/tmp/cache"
	serverConfig.System.Backup = "/Users/whoami/Documents/files/golang/sailhost.io/cloud/tmp/backup"
	serverConfig.System.EncryptKey = "cloud"

	global.CONF = serverConfig
	global.Viper = v
}
