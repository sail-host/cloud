package systemd

import "text/template"

var SystemdTemplate = template.Must(template.New("systemd").Parse(`[Unit]
Description={{ .Description }}

[Service]
Type={{ .Type }}
ExecStart={{ .ExecStart }}
WorkingDirectory={{ .WorkingDir }}
User={{ .User }}
{{range $key, $value := .Environment}}
Environment={{ $key }}={{ $value }}
{{end}}
Restart={{ .Restart }}
RestartSec=5s
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier={{ .SyslogIdentifier }}

[Install]
WantedBy=multi-user.target
`))
