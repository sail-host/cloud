package nodejs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

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
	nodePath := path.Join(nm.Path, fmt.Sprintf("nodejs/%s", nm.Version))

	// Check node binary file exist this path folder
	nodeBinary := path.Join(nodePath, "bin/node")
	if _, err := os.Stat(nodeBinary); os.IsNotExist(err) {
		return false, err
	}

	// Check Version this nodejs binary file
	version, err := exec.Command(nodeBinary, "--version").Output()
	if err != nil {
		return false, err
	}

	// TODO: Remove this line
	fmt.Println(string(version))

	// Check npm exists
	npmBinary := path.Join(nodePath, "bin/npm")
	if _, err := os.Stat(npmBinary); os.IsNotExist(err) {
		return false, err
	}

	// Check npx exists
	npxBinary := path.Join(nodePath, "bin/npx")
	if _, err := os.Stat(npxBinary); os.IsNotExist(err) {
		return false, err
	}

	// Check yarn exists
	yarnBinary := path.Join(nodePath, "bin/yarn")
	if _, err := os.Stat(yarnBinary); os.IsNotExist(err) {
		return false, err
	}

	// Check bun exists
	if _, err := os.Stat(path.Join(nodePath, "bin/bun")); os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}

func (nm *NodejsManager) InstallVersion(path string) error {
	// Check version exists
	exists, err := nm.CheckVersionExist()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// Create path for this version
	nodePath := filepath.Join(nm.Path, fmt.Sprintf("nodejs/%s", nm.Version))
	if err := os.MkdirAll(nodePath, 0755); err != nil {
		return err
	}

	// Download this version in the nodejs website
	downloadUrl := fmt.Sprintf("https://nodejs.org/dist/v%s/node-v%s-linux-x64.tar.xz", nm.Version, nm.Version)
	downloadFile := filepath.Join(nodePath, fmt.Sprintf("node-v%s-linux-x64.tar.xz", nm.Version))

	// Download file
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create file
	file, err := os.Create(downloadFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	// Unzip download file and move all files to correct path
	cmd := exec.Command("tar", "-xvf", downloadFile, "-C", nodePath)
	if err := cmd.Run(); err != nil {
		return err
	}

	// Import node binary to env PATH
	os.Setenv("PATH", fmt.Sprintf("%s:%s", os.Getenv("PATH"), nodePath))

	// Download or update npm and npx
	cmd = exec.Command("npm", "install", "-g", "npm@latest")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Download or update npx
	cmd = exec.Command("npm", "install", "-g", "npx@latest")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Download yarn and bun
	cmd = exec.Command("npm", "install", "-g", "yarn@latest")
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("npm", "install", "-g", "bun@latest")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Confirm installation. Check new version
	version, err := exec.Command(nodePath, "bin/node", "--version").Output()
	if err != nil {
		return err
	}

	// TODO: Remove this line
	fmt.Println(string(version))

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
