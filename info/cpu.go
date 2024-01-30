package info

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Create and show CPU info window
func CpuInfo(cpuVendor string, cpuModel string, app fyne.App) {

	fmt.Println(cpuVendor, "\n", cpuModel)

	cpuWindow := app.NewWindow("CPU Info")

	cpuData := [][]string{[]string{"CPU:", cpuModel},
		[]string{"Vendor:", cpuVendor}}

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

	backButton := widget.NewButton("Back", func() { cpuWindow.Close() })

	layout := container.New(layout.NewGridLayout(1), cpuTable, backButton)

	cpuWindow.SetContent(layout)
	cpuWindow.Resize(fyne.NewSize(380, 240))
	cpuWindow.CenterOnScreen()
	cpuWindow.Show()
}
