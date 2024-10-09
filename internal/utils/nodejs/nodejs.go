package nodejs

type NodejsManager struct {
	Version string
	Path    string
}

type INodejsManager interface {
	CheckVersionExist() (bool, error)
	InstallVersion(path string) error
	CmdNpmRun(command string) (string, error) // TODO: This return type string changed for correct return type!
	CmdBunRun(command string) (string, error)
}

func NewNodejsManager(version string, path string) INodejsManager {
	return &NodejsManager{
		Version: version,
		Path:    path,
	}
}

func (nm *NodejsManager) CheckVersionExist() (bool, error) {
	// Create this version path full

	// Check node binary file exist this path folder

	// Check Version this nodejs binary file

	// Check npm exists

	// Check npx exists

	// Check yarn exists

	// Check bun exists

	return true, nil
}

func (nm *NodejsManager) InstallVersion(path string) error {
	// Check version exists

	// Create path for this version

	// Download this version in the nodejs website

	// Unzip download file and move all files to correct path

	// Download or update npm and npx

	// Download yarn and bun

	// Confirm installation. Check new version

	return nil
}

func (nm *NodejsManager) CmdNpmRun(command string) (string, error) {
	// TODO: Implement this method

	return "", nil
}

func (nm *NodejsManager) CmdBunRun(command string) (string, error) {
	// TODO: Implement this method

	return "", nil
}
