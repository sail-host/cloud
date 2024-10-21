package config

import _ "embed"

//go:embed template.conf
var NginxTemplate []byte
