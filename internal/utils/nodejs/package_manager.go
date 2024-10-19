package nodejs

import (
	"errors"
	"os"
	"path/filepath"
)

type NodejsPackageManager struct {
	filesPath string
}

func NewNodejsPackageManager(filesPath string) *NodejsPackageManager {
	return &NodejsPackageManager{
		filesPath: filesPath,
	}
}

type PackageManager struct {
	manager []string
}

func (pm *NodejsPackageManager) Check() (*PackageManager, error) {
	var manager PackageManager

	// Check bun.lockb file for bun package manager
	if _, err := os.Stat(filepath.Join(pm.filesPath, "bun.lockb")); err == nil {
		manager.manager = append(manager.manager, "bun")
	}

	// Check pnpm-lock.yaml for pnpm package manager
	if _, err := os.Stat(filepath.Join(pm.filesPath, "pnpm-lock.yaml")); err == nil {
		manager.manager = append(manager.manager, "pnpm")
	}

	// Check yarn.lock file for yarn package manager
	if _, err := os.Stat(filepath.Join(pm.filesPath, "yarn.lock")); err == nil {
		manager.manager = append(manager.manager, "yarn")
	}

	// Check package.json file and set default npm for package manager
	if _, err := os.Stat(filepath.Join(pm.filesPath, "package.json")); err == nil {
		manager.manager = append(manager.manager, "npm")
	}

	if len(manager.manager) == 0 {
		return nil, errors.New("no package manager detected")
	}

	return &manager, nil
}
