package systemd

import (
	"os"
	"testing"
)

func TestIntegration_CreateAndManageService(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("Integration tests must be run as root")
	}

	manager := New()

	// Test NextJS service creation
	config := NextJSConfig{
		Port:        "3000",
		ProjectPath: "/tmp/test-nextjs",
	}
	service := NewNextJSService(config)

	// Create directory for testing
	os.MkdirAll("/tmp/test-nextjs", 0755)
	defer os.RemoveAll("/tmp/test-nextjs")

	// Test full lifecycle
	t.Run("service lifecycle", func(t *testing.T) {
		// Create
		if err := manager.CreateService(service); err != nil {
			t.Fatalf("Failed to create service: %v", err)
		}

		// Start
		if err := manager.StartService(service.Name); err != nil {
			t.Errorf("Failed to start service: %v", err)
		}

		// Restart
		if err := manager.RestartService(service.Name); err != nil {
			t.Errorf("Failed to restart service: %v", err)
		}

		// Stop
		if err := manager.StopService(service.Name); err != nil {
			t.Errorf("Failed to stop service: %v", err)
		}

		// Delete
		if err := manager.DeleteService(service.Name); err != nil {
			t.Errorf("Failed to delete service: %v", err)
		}
	})
}
