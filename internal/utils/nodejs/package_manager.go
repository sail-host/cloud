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

	// TODO: Impltement this method

	return &manager, nil
}
