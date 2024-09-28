package viper

import (
	"log"
	"os"
	"path"

	"github.com/sail-host/cloud/internal/config"
	"github.com/sail-host/cloud/internal/constants"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/random"
	"github.com/spf13/viper"
)

func Init(devMode bool) {
	global.CONF = createOrGetConfig(devMode)
}

// Config creator or get from file
func createOrGetConfig(devMode bool) config.Config {
	var rootConfig config.Config
	var configDir string
	var err error

	if devMode {
		configDir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		configDir = path.Join(configDir, "tmp")
	} else {
		configDir = constants.ConfigDir
	}

	// Ensure config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatalf("Failed to create config directory: %v", err)
	}

	file := path.Join(configDir, constants.ConfigFileName)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		// Config file does not exist, create it
		rootConfig = createConfig(devMode, configDir)
		viper.SetConfigFile(file)
		viper.SetConfigType("yaml")

		// Populate Viper with config values
		setViperConfig(rootConfig)

		// Write config to file
		if err := viper.WriteConfig(); err != nil {
			log.Fatalf("Failed to write config file: %v", err)
		}
	} else {
		// Config file exists, read it
		viper.SetConfigFile(file)
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}
		if err := viper.Unmarshal(&rootConfig); err != nil {
			log.Fatalf("Failed to unmarshal config: %v", err)
		}
	}

	return rootConfig
}

// Populate Viper with the config values
func setViperConfig(cfg config.Config) {
	viper.Set("debug", cfg.Debug)
	viper.Set("log.level", cfg.Log.Level)
	viper.Set("log.log_name", cfg.Log.LogName)
	viper.Set("log.log_suffix", cfg.Log.LogSuffix)
	viper.Set("log.max_backup", cfg.Log.MaxBackup)
	viper.Set("log.time_zone", cfg.Log.TimeZone)

	viper.Set("system.mode", cfg.System.Mode)
	viper.Set("system.ssl", cfg.System.SSL)
	viper.Set("system.backup", cfg.System.Backup)
	viper.Set("system.base_dir", cfg.System.BaseDir)
	viper.Set("system.bind_address", cfg.System.BindAddress)
	viper.Set("system.port", cfg.System.Port)
	viper.Set("system.cache", cfg.System.Cache)
	viper.Set("system.data_dir", cfg.System.DataDir)
	viper.Set("system.tmp_dir", cfg.System.TmpDir)
	viper.Set("system.db_file", cfg.System.DbFile)
	viper.Set("system.db_path", cfg.System.DbPath)
	viper.Set("system.log_path", cfg.System.LogPath)
	viper.Set("system.version", cfg.System.Version)
	viper.Set("system.encrypt_key", cfg.System.EncryptKey)
}

// Create config file
func createConfig(devMode bool, configDir string) config.Config {
	var rootConfig config.Config

	if devMode {
		rootConfig.Debug = true
		rootConfig.Log.Level = "debug"
		rootConfig.System.Mode = "dev"
	} else {
		rootConfig.Debug = false
		rootConfig.Log.Level = "info"
		rootConfig.System.Mode = "production"
	}

	rootConfig.System.SSL = ""
	rootConfig.Log.LogName = "cloud"
	rootConfig.Log.LogSuffix = ".log"
	rootConfig.Log.MaxBackup = 10
	rootConfig.Log.TimeZone = "Europe/Istanbul"

	rootConfig.System.Backup = path.Join(configDir, "backup")
	rootConfig.System.BaseDir = configDir
	rootConfig.System.BindAddress = "0.0.0.0"
	rootConfig.System.Port = "8080"
	rootConfig.System.Cache = path.Join(configDir, "cache")
	rootConfig.System.DataDir = path.Join(configDir, "data")
	rootConfig.System.TmpDir = path.Join(configDir, "tmp")
	rootConfig.System.DbFile = "cloud.db"
	rootConfig.System.DbPath = path.Join(configDir, "db")
	rootConfig.System.LogPath = path.Join(configDir, "log")
	rootConfig.System.Version = constants.Version
	rootConfig.System.EncryptKey = random.StringGenerator(32)

	return rootConfig
}
