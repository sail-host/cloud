package nginx

type NginxManager struct {
	sitesPath  string
	configPath string
}

type INginxManager interface {
	// TODO: Implement this methods
}

func NewNginxManager(sitesPath, configPath string) INginxManager {
	return &NginxManager{
		sitesPath:  sitesPath,
		configPath: configPath,
	}
}
