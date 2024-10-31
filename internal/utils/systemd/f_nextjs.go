package systemd

import (
	"path/filepath"
)

type NextJSConfig struct {
	Port         string
	ProjectPath  string
	StartCommand string
}

func NewNextJSService(config NextJSConfig) ServiceConfig {
	return ServiceConfig{
		Name:        filepath.Base(config.ProjectPath),
		Description: "NextJS Application Service",
		ExecStart:   config.StartCommand,
		WorkingDir:  config.ProjectPath,
		User:        "nodejs",
		Environment: map[string]string{
			"PORT": config.Port,
		},
		Restart:          "always",
		Type:             "simple",
		SyslogIdentifier: filepath.Base(config.ProjectPath),
	}
}
