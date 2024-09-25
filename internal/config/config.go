package config

type Config struct {
	System System    `mapstructure:"system"`
	Log    LogConfig `mapstructure:"log"`
	Debug  bool      `mapstructure:"debug"`
}
