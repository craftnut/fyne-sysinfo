package info

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var cpuList fyne.Widget

// Create and show CPU info window
func CpuInfo(
	cpuVendor string, cpuModel string,
	cpuCores string, cpuSpeed float64,
	app fyne.App,
) {

	cpuWindow := app.NewWindow("CPU Info")

	buildCpuList(
		cpuVendor, cpuModel,
		cpuCores, cpuSpeed,
	)

	layout := container.New(layout.NewStackLayout(), cpuList)

	cpuWindow.SetContent(layout)
	cpuWindow.Resize(fyne.NewSize(275, 225))
	cpuWindow.CenterOnScreen()
	cpuWindow.Show()

}

func buildCpuList(
	cpuVendor string, cpuModel string,
	cpuThreads string, cpuSpeed float64,
) {

	cpuModel = fmt.Sprintf("CPU: %s", cpuModel)
	cpuThreads = fmt.Sprintf("Threads: %s", cpuThreads)
	cpuSpeedStr := fmt.Sprintf("Speed: %f MHz", cpuSpeed)
	cpuVendor = fmt.Sprintf("Vendor: %s", cpuVendor)

	cpuListItems := []string{
		cpuModel,
		cpuThreads,
		cpuSpeedStr,
		cpuVendor,
	}

	cpuList = widget.NewList(
		func() int {
			return len(cpuListItems)
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(index int, obj fyne.CanvasObject) {
			if label, ok := obj.(*widget.Label); ok {
				label.SetText(cpuListItems[index])
			}
		},
	)
}
