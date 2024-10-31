package systemd

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

type mockSystemdManager struct {
	tempDir string
}

func newMockManager(t *testing.T) *mockSystemdManager {
	tempDir, err := os.MkdirTemp("", "systemd-test-")
	if err != nil {
		t.Fatal(err)
	}
	return &mockSystemdManager{tempDir: tempDir}
}

func (m *mockSystemdManager) cleanup() {
	os.RemoveAll(m.tempDir)
}

func TestSystemdManager_CreateService(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("Integration tests must be run as root")
	}
	if runtime.GOOS != "linux" {
		t.Skip("Integration tests must be run on Linux")
	}

	mock := newMockManager(t)
	defer mock.cleanup()

	manager := &SystemdManager{}

	config := ServiceConfig{
		Name:        "test-service",
		Description: "Test Service",
		ExecStart:   "/usr/bin/test",
		WorkingDir:  "/tmp",
		User:        "testuser",
		Environment: map[string]string{"TEST": "true"},
		Restart:     "always",
		Type:        "simple",
	}

	// Override systemd path for testing
	systemdPath = mock.tempDir

	err := manager.CreateService(config)
	if err != nil {
		t.Errorf("CreateService() error = %v", err)
	}

	// Check if service file was created
	servicePath := filepath.Join(mock.tempDir, config.Name+".service")
	if _, err := os.Stat(servicePath); os.IsNotExist(err) {
		t.Error("Service file was not created")
	}

	// Check file contents
	content, err := os.ReadFile(servicePath)
	if err != nil {
		t.Fatal(err)
	}

	// Verify content contains expected values
	expectedStrings := []string{
		"Description=Test Service",
		"ExecStart=/usr/bin/test",
		"User=testuser",
		"Environment=TEST=true",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(string(content), expected) {
			t.Errorf("Service file missing expected string: %s", expected)
		}
	}
}

func TestSystemdManager_DeleteService(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("Integration tests must be run as root")
	}
	if runtime.GOOS != "linux" {
		t.Skip("Integration tests must be run on Linux")
	}

	mock := newMockManager(t)
	defer mock.cleanup()

	manager := &SystemdManager{}
	systemdPath = mock.tempDir

	// Create a test service file
	serviceName := "test-service"
	servicePath := filepath.Join(mock.tempDir, serviceName+".service")
	if err := os.WriteFile(servicePath, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Test deletion
	if err := manager.DeleteService(serviceName); err != nil {
		t.Errorf("DeleteService() error = %v", err)
	}

	// Verify file was deleted
	if _, err := os.Stat(servicePath); !os.IsNotExist(err) {
		t.Error("Service file was not deleted")
	}
}
