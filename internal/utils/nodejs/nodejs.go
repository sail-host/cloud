package nodejs

type NodejsManager struct {
	Version string
}

type INodejsManager interface {
	CheckVersionExist() (bool, error)
	InstallVersion(path string) error
	CmdNpmRun(command string) (string, error)
}

func NewNodejsManager(version string) INodejsManager {
	return &NodejsManager{
		Version: version,
	}
}

func (nm *NodejsManager) CheckVersionExist() (bool, error) {
	// TODO: Implement this method

	return false, nil
}

func (nm *NodejsManager) InstallVersion(path string) error {
	// TODO: Implement this method

	return nil
}

func (nm *NodejsManager) CmdNpmRun(command string) (string, error) {
	// TODO: Implement this method

	return "", nil
}
