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

// TODO: Remove this line
// Nodejsda projectni commandlarni run qilishdan oldin nodejs pathni env PATHga qushib
// ishlatb bo'lgandan keyin uni pathdan olib tashlash kerak
// Nodejs uzini ham ishlatishigan oldin shundan qilsh kerak bo'ladi!
// Bu ishini barcha binary dasturlar uchun ishlatish mumkin.
// Projectda barcha qushimcha dasturlar bir folderda yig'ishga
// harakat qilish kerak. Dastur o'chirilgan vaqtida barchasini uchirish
// qulaylashadi. Ularning barchasi systemdagi dasturlardan bog'liqsiz
// ishlashi kerak.

func (nm *NodejsManager) CheckVersionExist() (bool, error) {
	// Create this version path full
	// nodePath := path.Join(nm.Path, fmt.Sprintf("nodejs/%s", nm.Version))

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
