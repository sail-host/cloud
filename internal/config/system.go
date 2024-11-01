package config

type System struct {
	Port        string `mapstructure:"port"`
	Ipv6        string `mapstructure:"ipv6"`
	BindAddress string `mapstructure:"bindAddress"`
	SSL         string `mapstructure:"ssl"`
	DbFile      string `mapstructure:"db_file"`
	DbPath      string `mapstructure:"db_path"`
	LogPath     string `mapstructure:"log_path"`
	DataDir     string `mapstructure:"data_dir"`
	TmpDir      string `mapstructure:"tmp_dir"`
	Cache       string `mapstructure:"cache"`
	Backup      string `mapstructure:"backup"`
	EncryptKey  string `mapstructure:"encrypt_key"`
	BaseDir     string `mapstructure:"base_dir"`
	Mode        string `mapstructure:"mode"`
	DeployDir   string `mapstructure:"deploy_dir"`
	UtilsDir    string `mapstructure:"utils_dir"`
}
