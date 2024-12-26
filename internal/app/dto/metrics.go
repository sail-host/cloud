package dto

type SystemMetrics struct {
	CPU  float64   `json:"cpu"`  // CPU usage percentage
	RAM  MemStats  `json:"ram"`  // RAM statistics
	Disk DiskStats `json:"disk"` // Disk statistics
	Time string    `json:"time"` // Timestamp
}

type MemStats struct {
	Total     uint64  `json:"total"`     // Total RAM
	Used      uint64  `json:"used"`      // Used RAM
	Available uint64  `json:"available"` // Available RAM
	Usage     float64 `json:"usage"`     // Usage percentage
}

type DiskStats struct {
	Total     uint64  `json:"total"`     // Total disk space
	Used      uint64  `json:"used"`      // Used disk space
	Available uint64  `json:"available"` // Available disk space
	Usage     float64 `json:"usage"`     // Usage percentage
}
