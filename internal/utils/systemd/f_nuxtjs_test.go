package systemd

import (
	"testing"
)

func TestNewNuxtJSService(t *testing.T) {
	tests := []struct {
		name   string
		config NuxtJSConfig
		want   ServiceConfig
	}{
		{
			name: "basic nuxtjs config",
			config: NuxtJSConfig{
				Port:         "3000",
				ProjectPath:  "/test/path",
				StartCommand: "npm start",
				ConfigName:   "path",
			},
			want: ServiceConfig{
				Name:        "path",
				Description: "NuxtJS Application Service",
				ExecStart:   "npm start",
				WorkingDir:  "/test/path",
				User:        "root",
				Environment: map[string]string{
					"PORT": "3000",
				},
				Restart:          "always",
				Type:             "simple",
				SyslogIdentifier: "path",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNuxtJSService(tt.config)
			if got.Name != tt.want.Name {
				t.Errorf("NewNuxtJSService().Name = %v, want %v", got.Name, tt.want.Name)
			}

			if got.ExecStart != tt.want.ExecStart {
				t.Errorf("NewNuxtJSService().ExecStart = %v, want %v", got.ExecStart, tt.want.ExecStart)
			}

			if got.WorkingDir != tt.want.WorkingDir {
				t.Errorf("NewNuxtJSService().WorkingDir = %v, want %v", got.WorkingDir, tt.want.WorkingDir)
			}

			if got.User != tt.want.User {
				t.Errorf("NewNuxtJSService().User = %v, want %v", got.User, tt.want.User)
			}
		})
	}
}
