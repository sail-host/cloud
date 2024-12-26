package service

import (
	"time"

	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type MetricsService struct{}

func NewMetricsService() *MetricsService {
	return &MetricsService{}
}

func (s *MetricsService) GetSystemMetrics() (*dto.SystemMetrics, error) {
	// Get CPU usage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	// Get memory stats
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// Get disk stats
	diskStats, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	return &dto.SystemMetrics{
		CPU: cpuPercent[0],
		RAM: dto.MemStats{
			Total:     memStats.Total,
			Used:      memStats.Used,
			Available: memStats.Available,
			Usage:     memStats.UsedPercent,
		},
		Disk: dto.DiskStats{
			Total:     diskStats.Total,
			Used:      diskStats.Used,
			Available: diskStats.Free,
			Usage:     diskStats.UsedPercent,
		},
		Time: time.Now().Format(time.RFC3339),
	}, nil
}
