package systemd

import (
	"fmt"
)

type ServiceConfig struct {
	Name             string
	Description      string
	ExecStart        string
	WorkingDir       string
	User             string
	Environment      map[string]string
	Restart          string
	Type             string
	SyslogIdentifier string
}

func (c ServiceConfig) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("service name cannot be empty")
	}
	if c.ExecStart == "" {
		return fmt.Errorf("ExecStart cannot be empty")
	}
	return nil
}

type ServiceManager interface {
	CreateService(config ServiceConfig) error
	UpdateService(name string, config ServiceConfig) error
	DeleteService(name string) error
	StartService(name string) error
	StopService(name string) error
	RestartService(name string) error
	ReloadService(name string) error
}
