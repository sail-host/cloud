package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func (m *Manager) ListDNSRecords(zoneID string, params cloudflare.ListDNSRecordsParams) ([]cloudflare.DNSRecord, error) {
	ctx := context.Background()
	records, _, err := m.client.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (m *Manager) CreateDNSRecord(zoneID string, params cloudflare.CreateDNSRecordParams) (cloudflare.DNSRecord, error) {
	ctx := context.Background()
	record, err := m.client.CreateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return cloudflare.DNSRecord{}, err
	}

	return record, nil
}

func (m *Manager) UpdateDNSRecord(zoneID string, params cloudflare.UpdateDNSRecordParams) (cloudflare.DNSRecord, error) {
	ctx := context.Background()
	record, err := m.client.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return cloudflare.DNSRecord{}, err
	}

	return record, nil
}

func (m *Manager) DeleteDNSRecord(zoneID string, recordID string) error {
	ctx := context.Background()
	err := m.client.DeleteDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), recordID)
	if err != nil {
		return err
	}

	return nil
}
