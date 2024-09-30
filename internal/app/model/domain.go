package model

type Domain struct {
	BaseModel
	Domain           string `json:"domain"`
	DNSProvider      string `json:"dns_provider" gorm:"enum('cloudflare', 'manual')"`
	CloudflareZoneID string `json:"cloudflare_zone_id" gorm:"type:varchar(255)"`
	CloudflareAPIKey string `json:"cloudflare_api_key" gorm:"type:varchar(255)"`
}
