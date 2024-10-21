package caddy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type SiteConfig struct {
	Domain string
	Root   string
	SSL    bool
	Proxy  string
}

type Caddy struct {
	Url string
}

func NewCaddy(url string) *Caddy {
	return &Caddy{Url: url}
}

func (c *Caddy) CreateSite(config *SiteConfig) error {
	listenPorts := []string{":80"}
	if config.SSL {
		listenPorts = append(listenPorts, ":443")
	}

	var routeHandler []map[string]interface{}
	if config.Proxy != "" {
		routeHandler = []map[string]interface{}{
			{
				"handler": "reverse_proxy",
				"upstreams": []map[string]interface{}{
					{
						"dial": config.Proxy,
					},
				},
			},
		}
	} else {
		routeHandler = []map[string]interface{}{
			{
				"handler": "file_server",
				"root":    config.Root,
			},
		}
	}

	caddyConfig := map[string]interface{}{
		"apps": map[string]interface{}{
			"http": map[string]interface{}{
				"servers": map[string]interface{}{
					strings.Replace(config.Domain, ".", "_", -1): map[string]interface{}{
						"listen": listenPorts,
						"routes": []map[string]interface{}{
							{
								"match": []map[string]interface{}{
									{
										"host": []string{config.Domain},
									},
								},
								"handle": routeHandler,
							},
						},
					},
				},
			},
		},
	}

	if config.SSL {
		caddyConfig["apps"].(map[string]interface{})["http"].(map[string]interface{})["servers"].(map[string]interface{})[strings.Replace(config.Domain, ".", "_", -1)].(map[string]interface{})["tls_connection_policies"] = []map[string]interface{}{
			{
				"match": map[string]interface{}{
					"sni": []string{config.Domain},
				},
			},
		}
	}

	jsonData, err := json.Marshal(caddyConfig)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	apiURL := fmt.Sprintf("http://%s/load", c.Url)

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending request to Caddy API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Caddy API responded with status: %s", resp.Status)
	}

	return nil
}

func (c *Caddy) RemoveSite(domain string) error {
	serverName := strings.Replace(domain, ".", "_", -1)
	apiURL := fmt.Sprintf("http://%s/config/apps/http/servers/%s", c.Url, serverName)

	req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
	if err != nil {
		return fmt.Errorf("error creating DELETE request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending DELETE request to Caddy API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Caddy API responded with status: %s", resp.Status)
	}

	return nil
}

func (c *Caddy) UpdateSite(config *SiteConfig) error {
	listenPorts := []string{":80"}
	if config.SSL {
		listenPorts = append(listenPorts, ":443")
	}

	var routeHandler []map[string]interface{}
	if config.Proxy != "" {
		routeHandler = []map[string]interface{}{
			{
				"handler": "reverse_proxy",
				"upstreams": []map[string]interface{}{
					{
						"dial": config.Proxy,
					},
				},
			},
		}
	} else {
		routeHandler = []map[string]interface{}{
			{
				"handler": "file_server",
				"root":    config.Root,
			},
		}
	}

	caddyConfig := map[string]interface{}{
		"apps": map[string]interface{}{
			"http": map[string]interface{}{
				"servers": map[string]interface{}{
					strings.Replace(config.Domain, ".", "_", -1): map[string]interface{}{
						"listen": listenPorts,
						"routes": []map[string]interface{}{
							{
								"match": []map[string]interface{}{
									{
										"host": []string{config.Domain},
									},
								},
								"handle": routeHandler,
							},
						},
					},
				},
			},
		},
	}

	if config.SSL {
		caddyConfig["apps"].(map[string]interface{})["http"].(map[string]interface{})["servers"].(map[string]interface{})[strings.Replace(config.Domain, ".", "_", -1)].(map[string]interface{})["tls_connection_policies"] = []map[string]interface{}{
			{
				"match": map[string]interface{}{
					"sni": []string{config.Domain},
				},
			},
		}
	}

	jsonData, err := json.Marshal(caddyConfig)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	apiURL := fmt.Sprintf("http://%s/load", c.Url)

	req, err := http.NewRequest(http.MethodPatch, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating PATCH request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending PATCH request to Caddy API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Caddy API responded with status: %s", resp.Status)
	}

	return nil
}
