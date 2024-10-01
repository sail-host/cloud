package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
)

type Manager struct {
	client *cloudflare.API
}

func NewManager(apiToken string) (*Manager, error) {
	client, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		return nil, err
	}

	return &Manager{client: client}, nil
}
