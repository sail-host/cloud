package dto

import "time"

type Domain struct {
	ID          uint      `json:"id"`
	Domain      string    `json:"domain"`
	DNSProvider string    `json:"dns_provider"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type DomainListResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []Domain `json:"data"`
}

type CreateDomainRequest struct {
	Domain           string `json:"domain" validate:"required"`
	DNSProvider      string `json:"dns_provider" validate:"required"`
	CloudflareZoneID string `json:"cloudflare_zone_id"`
	CloudflareAPIKey string `json:"cloudflare_api_key"`
}

type UpdateDomainRequest struct {
	Domain           string `json:"domain" validate:"required"`
	DNSProvider      string `json:"dns_provider" validate:"required"`
	CloudflareZoneID string `json:"cloudflare_zone_id"`
	CloudflareAPIKey string `json:"cloudflare_api_key"`
}

type CheckDomainRequest struct {
	Domain           string `json:"domain" validate:"required"`
	CloudflareZoneID string `json:"cloudflare_zone_id" validate:"required"`
	CloudflareAPIKey string `json:"cloudflare_api_key" validate:"required"`
}
