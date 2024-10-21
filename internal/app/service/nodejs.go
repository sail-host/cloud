package service

import (
	"path"

	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/nodejs"
)

type NodejsDeploymentService struct{}

type INodejsDeploymentService interface {
	InstallDependencies(deployment *model.Deployment) error
	Build(deployment *model.Deployment) error
}

func NewINodejsDeploymentService() INodejsDeploymentService {
	return &NodejsDeploymentService{}
}

func (n *NodejsDeploymentService) InstallDependencies(deployment *model.Deployment) error {
	// Check if node is installed
	nodeManager := nodejs.NewNodejsManager("v20", global.CONF.System.UtilsDir)

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

	project, err := projectRepo.GetProjectByID(deployment.ProjectID)
	if err != nil {
		global.LOG.Error("Error getting project", err)
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
