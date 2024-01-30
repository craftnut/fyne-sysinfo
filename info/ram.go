package info

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Create and show RAM window
func RamInfo(totalRam int, avalableRam int, usedRam int, usedPercent float64, app fyne.App) {

	ramWindow := app.NewWindow("RAM Info")

	ramData := [][]string{{"Total:", strconv.Itoa(totalRam) + " GB"},
		{"Used:", strconv.Itoa(usedRam) + " GB, " + fmt.Sprintf("%f", usedPercent) + "%"},
		{"Available:", strconv.Itoa(avalableRam) + " GB"}}

	ramTable := widget.NewTable(
		func() (rows int, cols int) {
			return len(ramData), len(ramData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("ram table")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ramData[i.Row][i.Col])
		},
	)

	layout := container.New(layout.NewStackLayout(), ramTable)

	ramWindow.SetContent(layout)
	ramWindow.Resize(fyne.NewSize(380, 240))
	ramWindow.CenterOnScreen()
	ramWindow.Show()
}
