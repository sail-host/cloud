package systemd

import "path/filepath"

type NuxtJSConfig struct {
	Port         string
	ProjectPath  string
	StartCommand string
}

func NewNuxtJSService(config NuxtJSConfig) ServiceConfig {
	return ServiceConfig{
		Name:        filepath.Base(config.ProjectPath),
		Description: "NuxtJS Application Service",
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
