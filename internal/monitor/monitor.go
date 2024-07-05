package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemMonitor struct {
	Fps            float64
	FrameCount     uint64
	LastFrameTime  time.Time
	LastUpdateTime time.Time
}

func NewSystemMonitor() *SystemMonitor {
	return &SystemMonitor{
		LastFrameTime:  time.Now(),
		LastUpdateTime: time.Now(),
	}
}

func (sm *SystemMonitor) UpdateFPS() {
	sm.FrameCount++
	now := time.Now()
	duration := now.Sub(sm.LastFrameTime).Seconds()

	if duration >= 1.0 {
		sm.Fps = float64(sm.FrameCount) / duration
		sm.FrameCount = 0
		sm.LastFrameTime = now
	}
}

func (sm *SystemMonitor) GetCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

func (sm *SystemMonitor) GetMemoryUsage() (float64, uint64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return vmStat.UsedPercent, vmStat.Total, nil
}

var Monitor = NewSystemMonitor()

func (sm *SystemMonitor) Refresh() {
	for {
		cpuUsage, err := sm.GetCPUUsage()
		if err != nil {
			fmt.Printf("Error getting CPU usage: %v\n", err)
			continue
		}

		memoryUsagePercent, totalMemory, err := sm.GetMemoryUsage()
		if err != nil {
			fmt.Printf("Error getting memory usage: %v\n", err)
			continue
		}

		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
		fmt.Printf("Memory Usage: %.2f%% of %d KB\n", memoryUsagePercent, totalMemory/1024)
		fmt.Printf("FPS: %.2f\n", sm.Fps)

		time.Sleep(time.Second)
	}
}
