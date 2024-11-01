package systemd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var systemdPath = "/etc/systemd/system"

type SystemdManager struct{}

func New() *SystemdManager {
	return &SystemdManager{}
}

func (m *SystemdManager) CreateService(config ServiceConfig) error {
	servicePath := filepath.Join(systemdPath, config.Name+".service")

	file, err := os.Create(servicePath)
	if err != nil {
		return fmt.Errorf("failed to create service file: %w", err)
	}
	defer file.Close()

	if err := SystemdTemplate.Execute(file, config); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}

	if err = m.reloadDaemon(); err != nil {
		return fmt.Errorf("failed to reload systemd daemon: %w", err)
	}

	if err = m.EnableService(config.Name); err != nil {
		return fmt.Errorf("failed to enable service: %w", err)
	}

	return nil
}

func (m *SystemdManager) UpdateService(name string, config ServiceConfig) error {
	if err := m.DeleteService(name); err != nil {
		return err
	}
	return m.CreateService(config)
}

func (m *SystemdManager) DeleteService(name string) error {
	servicePath := filepath.Join(systemdPath, name+".service")

	if err := m.DisableService(name); err != nil {
		return fmt.Errorf("failed to disable service: %w", err)
	}

	if err := os.Remove(servicePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete service file: %w", err)
	}
	return m.reloadDaemon()
}

func (m *SystemdManager) StartService(name string) error {
	return m.runSystemctl("start", name)
}

func (m *SystemdManager) StopService(name string) error {
	return m.runSystemctl("stop", name)
}

func (m *SystemdManager) RestartService(name string) error {
	return m.runSystemctl("restart", name)
}

func (m *SystemdManager) ReloadService(name string) error {
	return m.runSystemctl("reload", name)
}

func (m *SystemdManager) reloadDaemon() error {
	return m.runSystemctl("daemon-reload", "")
}

func (m *SystemdManager) EnableService(name string) error {
	return m.runSystemctl("enable", name)
}

func (m *SystemdManager) DisableService(name string) error {
	return m.runSystemctl("disable", name)
}

func (m *SystemdManager) runSystemctl(command, serviceName string) error {
	args := []string{command}
	if serviceName != "" {
		args = append(args, serviceName+".service")
	}

	cmd := exec.Command("systemctl", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("systemctl %s failed: %w", command, err)
	}
	return nil
}
