package service

import (
	"fmt"
	"path"

	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/nodejs"
	"github.com/sail-host/cloud/internal/utils/systemd"
)

type NodejsDeploymentService struct{}

type INodejsDeploymentService interface {
	InstallDependencies(deployment *model.Deployment) error
	Build(deployment *model.Deployment) error
	Start(deployment *model.Deployment) error
}

func NewINodejsDeploymentService() INodejsDeploymentService {
	return &NodejsDeploymentService{}
}

func (n *NodejsDeploymentService) InstallDependencies(deployment *model.Deployment) error {
	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		global.LOG.Error("Error getting project", err)
		return err
	}

	// Check if node is installed
	nodeManager := nodejs.NewNodejsManager(project.NodeVersion, global.CONF.System.UtilsDir)

	exists, err := nodeManager.CheckVersionExist()
	if err != nil {
		global.LOG.Error("Error checking Nodejs version", err)
		return err
	}
	if !exists {
		err = projectRepo.CreateLog(deployment, "Nodejs version not exists. Installing...")
		if err != nil {
			global.LOG.Error("Error creating deployment log", err)
			return err
		}
		err = nodeManager.InstallVersion()
		if err != nil {
			global.LOG.Error("Error installing Nodejs", err)
			return err
		}
	}
	runPath := path.Join(global.CONF.System.DeployDir, deployment.UUID)

	packageManager := nodejs.NewNodejsPackageManager(runPath)
	managers, err := packageManager.Check()
	if err != nil {
		global.LOG.Error("Error checking nodejs package manager", err)
		return err
	}

	// Check user write install command
	if project.InstallCommand != "" {
		_, err = nodeManager.Bash(project.InstallCommand, runPath)
		if err != nil {
			global.LOG.Error("Error running install command", err)
			return err
		}
	} else {
		// Check nodejs package manager

		switch managers.Manager[0] {
		case "bun":
			_, err = nodeManager.CmdBunRun("install", runPath)
		case "pnpm":
			_, err = nodeManager.CmdPnpmRun("install", runPath)
		case "yarn":
			_, err = nodeManager.CmdYarnRun("install", runPath)
		default:
			_, err = nodeManager.CmdNpmRun("install", runPath)
		}
		if err != nil {
			global.LOG.Error("Error installing nodejs package manager", err)
			return err
		}
	}

	return nil
}

func (n *NodejsDeploymentService) Build(deployment *model.Deployment) error {
	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		global.LOG.Error("Error getting project", err)
		return err
	}

	nodeManager := nodejs.NewNodejsManager("v20", global.CONF.System.UtilsDir)
	runPath := path.Join(global.CONF.System.DeployDir, deployment.UUID)

	if project.BuildCommand != "" {
		_, err = nodeManager.CmdNpmRun(project.BuildCommand, runPath)
		if err != nil {
			global.LOG.Error("Error running build command", err)
			return err
		}
	} else {
		packageManager := nodejs.NewNodejsPackageManager(runPath)
		managers, err := packageManager.Check()
		if err != nil {
			global.LOG.Error("Error checking nodejs package manager", err)
			return err
		}
		switch managers.Manager[0] {
		case "bun":
			_, err = nodeManager.CmdBunRun("run build", runPath)
		case "pnpm":
			_, err = nodeManager.CmdPnpmRun("run build", runPath)
		case "yarn":
			_, err = nodeManager.CmdYarnRun("build", runPath)
		default:
			_, err = nodeManager.CmdNpmRun("run build", runPath)
		}
		if err != nil {
			global.LOG.Error("Error running build command", err)
			return err
		}
	}

	return nil
}

func (n *NodejsDeploymentService) Start(deployment *model.Deployment) error {
	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		global.LOG.Error("Error getting project", err)
		return err
	}

	manager := systemd.New()
	configName := fmt.Sprintf("sailhost-%d", deployment.ID)
	port := fmt.Sprintf("%d", 10000+deployment.ID)

	if project.Framework == "nextjs" {
		service := systemd.NewNextJSService(systemd.NextJSConfig{
			Port:         port,
			ProjectPath:  project.RootDir, // TODO: check if this is correct
			StartCommand: "npm run start", // TODO: check if this is correct
			ConfigName:   configName,
		})

		err := manager.CreateService(service)
		if err != nil {
			global.LOG.Error("Error creating NextJS service", err)
			return err
		}
	}

	if project.Framework == "nuxtjs" {
		service := systemd.NewNuxtJSService(systemd.NuxtJSConfig{
			Port:         port,
			ProjectPath:  project.RootDir,
			StartCommand: "npm run start",
			ConfigName:   configName,
		})

		err := manager.CreateService(service)
		if err != nil {
			global.LOG.Error("Error creating NuxtJS service", err)
			return err
		}
	}

	return nil
}
