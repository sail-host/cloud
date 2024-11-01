package systemd

import "testing"

func TestNewNextJSService(t *testing.T) {
	tests := []struct {
		name   string
		config NextJSConfig
		want   ServiceConfig
	}{
		{
			name: "basic nextjs config",
			config: NextJSConfig{
				Port:         "3000",
				ProjectPath:  "/test/path",
				StartCommand: "npm start",
				ConfigName:   "path",
			},
			want: ServiceConfig{
				Name:        "path",
				Description: "NextJS Application Service",
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
			got := NewNextJSService(tt.config)
			if got.Name != tt.want.Name {
				t.Errorf("NewNextJSService().Name = %v, want %v", got.Name, tt.want.Name)
			}

			if got.ExecStart != tt.want.ExecStart {
				t.Errorf("NewNextJSService().ExecStart = %v, want %v", got.ExecStart, tt.want.ExecStart)
			}

			if got.WorkingDir != tt.want.WorkingDir {
				t.Errorf("NewNextJSService().WorkingDir = %v, want %v", got.WorkingDir, tt.want.WorkingDir)
			}

			if got.User != tt.want.User {
				t.Errorf("NewNextJSService().User = %v, want %v", got.User, tt.want.User)
			}
		})
	}
}
