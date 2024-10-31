package systemd

type NuxtJSConfig struct {
	Port         string
	ProjectPath  string
	StartCommand string
	ConfigName   string
}

func NewNuxtJSService(config NuxtJSConfig) ServiceConfig {
	return ServiceConfig{
		Name:        config.ConfigName,
		Description: "NuxtJS Application Service",
		ExecStart:   config.StartCommand,
		WorkingDir:  config.ProjectPath,
		User:        "root",
		Environment: map[string]string{
			"PORT": config.Port,
		},
		Restart:          "always",
		Type:             "simple",
		SyslogIdentifier: config.ConfigName,
	}
}
