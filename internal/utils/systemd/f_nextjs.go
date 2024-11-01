package systemd

type NextJSConfig struct {
	Port         string
	ProjectPath  string
	StartCommand string
	ConfigName   string
}

func NewNextJSService(config NextJSConfig) ServiceConfig {
	return ServiceConfig{
		Name:        config.ConfigName,
		Description: "NextJS Application Service",
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
