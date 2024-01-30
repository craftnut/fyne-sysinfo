package info

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Create and show CPU info window
func CpuInfo(cpuVendor string, cpuModel string, cpuCores int32, cpuSpeed float64, app fyne.App) {

	fmt.Println(cpuVendor, "\n", cpuModel)

	cpuWindow := app.NewWindow("CPU Info")

	cpuData := [][]string{{"CPU:", cpuModel},
		{"Cores:", strconv.Itoa(int(cpuCores))},
		{"Speed:", fmt.Sprintf("%f", cpuSpeed) + " MHz"},
		{"Vendor:", cpuVendor}}

	cpuTable := widget.NewTable(
		func() (rows int, cols int) {
			return len(cpuData), len(cpuData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("cpu table")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(cpuData[i.Row][i.Col])
		},
	)

	layout := container.New(layout.NewStackLayout(), cpuTable)

	cpuWindow.SetContent(layout)
	cpuWindow.Resize(fyne.NewSize(380, 240))
	cpuWindow.CenterOnScreen()
	cpuWindow.Show()
}
