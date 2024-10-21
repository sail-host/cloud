package nodejs

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sail-host/cloud/internal/global"
)

const NODEJS_API_URL = "https://nodejs.org/dist/index.json"

type NodejsManager struct {
	Version string
	Path    string
}

type INodejsManager interface {
	CheckVersionExist() (bool, error)
	InstallVersion() error
	CmdNpmRun(command string, runPath string) (string, error)
	CmdBunRun(command string, runPath string) (string, error)
	CmdYarnRun(command string, runPath string) (string, error)
	CmdPnpmRun(command string, runPath string) (string, error)
	Bash(command string, runPath string) (string, error)
}

func NewNodejsManager(version string, utilsPath string) INodejsManager {
	return &NodejsManager{
		Version: version,
		Path:    utilsPath,
	}
}

func (nm *NodejsManager) CheckVersionExist() (bool, error) {
	version, err := nm.Bash("node --version", "")
	if err != nil {
		return false, nil
	}
	global.LOG.Info("Nodejs version", version)

	_, err = nm.Bash("npm --version", "")
	if err != nil {
		return false, nil
	}

	_, err = nm.Bash("bun --version", "")
	if err != nil {
		return false, nil
	}

	_, err = nm.Bash("yarn --version", "")
	if err != nil {
		return false, nil
	}

	_, err = nm.Bash("pnpm --version", "")
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (nm *NodejsManager) InstallVersion() error {
	// Check version exists
	exists, err := nm.CheckVersionExist()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	global.LOG.Info("Installing Nodejs version", nm.Version)

	// Create path for this version
	nodePath := filepath.Join(nm.Path, fmt.Sprintf("nodejs/%s", nm.Version))
	if err := os.MkdirAll(nodePath, 0755); err != nil {
		return err
	}

	// Get os and arch
	runOS := runtime.GOOS
	runArch := runtime.GOARCH

	if runOS == "darwin" && runArch == "arm64" {
		runOS = "darwin"
		runArch = "arm64"
	} else if runOS == "darwin" {
		runArch = "x64"
	}

	version, err := getNodePath(nm.Version)
	if err != nil {
		global.LOG.Error("Error getting Nodejs version", err)
		return err
	}

	downloadURL := getDownloadURL(version, runOS, runArch)
	zipFilePath := fmt.Sprintf("node-%s-%s-%s.zip", version, runOS, runArch)

	// Download file
	if err := downloadFile(downloadURL, zipFilePath); err != nil {
		global.LOG.Error("Error downloading Nodejs file", err)
		return err
	}

	// Unzip file
	if err := unzipFile(zipFilePath, nodePath); err != nil {
		global.LOG.Error("Error unzipping Nodejs file", err)
		return err
	}

	// Remove zip file
	if err := os.Remove(zipFilePath); err != nil {
		global.LOG.Error("Error removing Nodejs zip file", err)
	}

	// Install or update npm
	_, err = nm.Bash("npm install --global npm@latest", "")
	if err != nil {
		global.LOG.Error("Error install or update npm :", err)
		return err
	}

	// Install yarn
	_, err = nm.Bash("npm install --global yarn", "")
	if err != nil {
		global.LOG.Error("Error install yarn :", err)
		return err
	}

	// Install bun
	_, err = nm.Bash("npm install --global bun", "")
	if err != nil {
		global.LOG.Error("Error install bun :", err)
		return err
	}

	// Install pnpm
	_, err = nm.Bash("npm install --global pnpm", "")
	if err != nil {
		global.LOG.Error("Error install pnpm :", err)
		return err
	}

	// Confirm installation. Check new version
	versionBytes, err := nm.Bash("node --version", "")
	if err != nil {
		global.LOG.Error("Error check nodejs version", err)
		return err
	}

	global.LOG.Info("Nodejs version", versionBytes)

	return nil
}

func (nm *NodejsManager) Bash(command string, runPath string) (string, error) {
	nodePath := filepath.Join(nm.Path, fmt.Sprintf("nodejs/%s/bin", nm.Version))

	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Dir = runPath
	cmd.Env = append(os.Environ(), fmt.Sprintf("PATH=%s:%s", nodePath, os.Getenv("PATH")))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command failed: %v\nStderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

// Parse nodejs versions and return last version
func getNodePath(version string) (string, error) {
	resp, err := http.Get(NODEJS_API_URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	type NodejsApiResponse struct {
		Version string `json:"version"`
	}

	var response []NodejsApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	var latestVersion string
	for _, v := range response {
		if len(v.Version) > 0 && v.Version[0:len(version)] == version {
			latestVersion = v.Version
			break
		}
	}

	if len(latestVersion) == 0 {
		return "", fmt.Errorf("version not found")
	}

	return latestVersion, nil
}

// Get Nodejs download url
func getDownloadURL(version, os, arch string) string {
	fileFormat := "tar.gz"
	if os == "windows" {
		fileFormat = "zip"
	}

	return fmt.Sprintf("https://nodejs.org/dist/%s/node-%s-%s-%s.%s", version, version, os, arch, fileFormat)
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Unzip downloaded nodejs zip file
func unzipFile(filepath string, destPath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	var baseFolder string

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if baseFolder == "" {
			baseFolder = strings.Split(header.Name, "/")[0]
		}

		relativePath := strings.TrimPrefix(header.Name, baseFolder+"/")
		target := path.Join(destPath, relativePath)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, header.FileInfo().Mode()); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, header.FileInfo().Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tr); err != nil {
				return err
			}
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, target); err != nil {
				return err
			}
		}
	}

	return nil
}

func (nm *NodejsManager) CmdNpmRun(command string, runPath string) (string, error) {
	return nm.Bash(fmt.Sprintf("npm %s", command), runPath)
}

func (nm *NodejsManager) CmdBunRun(command string, runPath string) (string, error) {
	return nm.Bash(fmt.Sprintf("bun %s", command), runPath)
}

func (nm *NodejsManager) CmdYarnRun(command string, runPath string) (string, error) {
	return nm.Bash(fmt.Sprintf("yarn %s", command), runPath)
}

func (nm *NodejsManager) CmdPnpmRun(command string, runPath string) (string, error) {
	return nm.Bash(fmt.Sprintf("pnpm %s", command), runPath)
}
