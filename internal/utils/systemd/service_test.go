package systemd

import "testing"

func TestServiceConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  ServiceConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: ServiceConfig{
				Name:             "test-service",
				Description:      "Test Service",
				ExecStart:        "/usr/bin/test",
				WorkingDir:       "/tmp",
				User:             "testuser",
				Environment:      map[string]string{"TEST": "true"},
				Restart:          "always",
				Type:             "simple",
				SyslogIdentifier: "test-service",
			},
			wantErr: false,
		},
		{
			name: "invalid empty name",
			config: ServiceConfig{
				Name:      "",
				ExecStart: "/usr/bin/test",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.config.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ServiceConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
