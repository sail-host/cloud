package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Port  string `env:"PORT" envDefault:"2050"`
	Env   string `env:"ENV" envDefault:"production"`
	Debug bool   `env:"DEBUG" envDefault:"false"`
}

var cfg Config

func GetConfig() *Config {
	if cfg == (Config{}) {
		cfg = Config{}
	}
	return &cfg
}

func LoadConfig() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}
}
