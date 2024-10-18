package nodejs

// TODO: Add this file check auto package managers in nodejs.

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

	// Check pnpm-lock.yaml for pnpm package manager

	// Check yarn.lock file for yarn package manager

	// Check packege.json file and set default npm for package manager

	return &manager, nil
}
