package nodejs

import (
	"os"
	"path/filepath"
)

func GetVersion(projectPath string) (string, error) {
	versionFile := filepath.Join(projectPath, ".node-version")

	// Check if file exists
	if _, err := os.Stat(versionFile); os.IsNotExist(err) {
		return "", nil
	}

	versionBytes, err := os.ReadFile(versionFile)
	if err != nil {
		return "", err
	}

	version := string(versionBytes)
	if len(version) < 1 || version[0] != 'v' {
		return "", nil
	}

	// Check if version format is like v16, v18, v20 etc
	if len(version) < 2 || !isNumeric(version[1]) {
		return "", nil
	}

	return version[:3], nil
}

func isNumeric(c byte) bool {
	return c >= '0' && c <= '9'
}
