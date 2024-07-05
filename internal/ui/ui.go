package ui

import (
	"fmt"
	"time"

	"go_sys_monitor/internal/monitor"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {
	a := app.New()
	w := a.NewWindow("System Monitor")

	cpuUsageLabel := widget.NewLabel("CPU Usage: ")
	memoryUsageLabel := widget.NewLabel("Memory Usage: ")
	fpsLabel := widget.NewLabel("FPS: ")

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			cpuUsage, err := monitor.Monitor.GetCPUUsage()
			if err != nil {
				cpuUsageLabel.SetText("Error getting CPU usage")
				continue
			}

			memoryUsagePercent, totalMemory, err := monitor.Monitor.GetMemoryUsage()
			if err != nil {
				memoryUsageLabel.SetText("Error getting memory usage")
				continue
			}

			cpuUsageLabel.SetText(fmt.Sprintf("CPU Usage: %.2f%%", cpuUsage))
			memoryUsageLabel.SetText(fmt.Sprintf("Memory Usage: %.2f%% of %d KB", memoryUsagePercent, totalMemory/1024))
			fpsLabel.SetText(fmt.Sprintf("FPS: %.2f", monitor.Monitor.Fps))
		}
	}()

	w.SetContent(container.NewVBox(
		cpuUsageLabel,
		memoryUsageLabel,
		fpsLabel,
	))

	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
