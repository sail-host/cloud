package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type DeploymentSystemdService struct {
}

var systemdTemplate = `[Unit]
Description={{.Description}}
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory={{.WorkingDirectory}}
Environment=PORT={{.Port}}
ExecStart={{.ExecStart}}
Restart=always
RestartSec=10   # Restart service after 10 seconds if node service crashes
StandardOutput=syslog  # Output to syslog
StandardError=syslog   # Output errors to syslog
SyslogIdentifier={{.SyslogIdentifier}}

[Install]
WantedBy=multi-user.target`

type IDeploymentSystemdService interface {
	CreateConfig(deployment *model.Deployment, workingDirectory string) error
	RestartService(deployment *model.Deployment, workingDirectory string) error
}

func NewIDeploymentSystemdService() IDeploymentSystemdService {
	return &DeploymentSystemdService{}
}

func (s *DeploymentSystemdService) CreateConfig(deployment *model.Deployment, workingDirectory string) error {
	if global.CONF.System.Mode == "dev" {
		return nil
	}

	// Check project framework have systemd service
	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		return err
	}

	// Create systemd service
	var run string
	switch project.Framework {
	case "nextjs":
		run = "npm run start"
	case "nuxt":
		run = "node server/index.mjs"
	default:
		return nil
	}

	port := 10000 + project.ID

	// Convert template to string
	tmpl, err := template.New("systemd").Parse(systemdTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, map[string]string{
		"Description":      fmt.Sprintf("%s - %s", project.Name, "Sailhost systemd"),
		"WorkingDirectory": workingDirectory,
		"Port":             fmt.Sprintf("%d", port),
		"ExecStart":        run,
		"SyslogIdentifier": fmt.Sprintf("%s-%d", project.Name, deployment.ID),
	}); err != nil {
		return err
	}

	filePath := fmt.Sprintf("/etc/systemd/system/sailhost-%d.service", project.ID)
	if err := os.WriteFile(filePath, buf.Bytes(), 0644); err != nil {
		return err
	}

	// Reload systemd
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return err
	}

	// Enable service
	if err := exec.Command("systemctl", "enable", fmt.Sprintf("sailhost-%d.service", project.ID)).Run(); err != nil {
		return err
	}

	// Start service
	if err := exec.Command("systemctl", "start", fmt.Sprintf("sailhost-%d.service", project.ID)).Run(); err != nil {
		return err
	}

	return nil
}

func (s *DeploymentSystemdService) RestartService(deployment *model.Deployment, workingDirectory string) error {
	if global.CONF.System.Mode == "dev" {
		return nil
	}

	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		return err
	}

	// Create systemd service
	var run string
	switch project.Framework {
	case "nextjs":
		run = "npm run start"
	case "nuxt":
		run = "node server/index.mjs"
	default:
		return nil
	}

	port := 10000 + project.ID

	// Convert template to string
	tmpl, err := template.New("systemd").Parse(systemdTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, map[string]string{
		"Description":      fmt.Sprintf("%s - %s", project.Name, "Sailhost systemd"),
		"WorkingDirectory": workingDirectory,
		"Port":             fmt.Sprintf("%d", port),
		"ExecStart":        run,
		"SyslogIdentifier": fmt.Sprintf("%s-%d", project.Name, deployment.ID),
	}); err != nil {
		return err
	}

	filePath := fmt.Sprintf("/etc/systemd/system/sailhost-%d.service", project.ID)
	// Update file
	if err := os.WriteFile(filePath, buf.Bytes(), 0644); err != nil {
		return err
	}

	// Reload systemd
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return err
	}

	// Restart service
	if err := exec.Command("systemctl", "restart", fmt.Sprintf("sailhost-%d.service", project.ID)).Run(); err != nil {
		return err
	}

	return nil
}
