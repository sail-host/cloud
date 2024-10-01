package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func (m *Manager) ListZones() ([]cloudflare.Zone, error) {
	ctx := context.Background()
	zones, err := m.client.ListZones(ctx)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (m *Manager) CheckZoneID(zoneID string) (bool, error) {
	list, err := m.ListZones()
	if err != nil {
		return false, err
	}

	for _, zone := range list {
		if zone.ID == zoneID {
			return true, nil
		}
	}

	return false, nil
}

func (m *Manager) ZoneInfo(zoneID string) (cloudflare.Zone, error) {
	var zone cloudflare.Zone
	list, err := m.ListZones()
	if err != nil {
		return zone, err
	}

	for _, zone := range list {
		if zone.ID == zoneID {
			return zone, nil
		}
	}

	return zone, nil
}
