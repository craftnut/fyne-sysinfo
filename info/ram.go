package info

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var ramList fyne.Widget

// Create and show RAM window
func RamInfo(
	totalRam int, availableRam int,
	usedRam int, usedPercent float64,
	app fyne.App,
) {

	ramWindow := app.NewWindow("RAM Info")

	buildRamList(
		totalRam, availableRam,
		usedRam, usedPercent,
	)

	layout := container.New(layout.NewStackLayout(), ramList)

	ramWindow.SetContent(layout)
	ramWindow.Resize(fyne.NewSize(275, 225))
	ramWindow.CenterOnScreen()
	ramWindow.Show()

}

func buildRamList(
	totalRam int, availableRam int,
	usedRam int, usedPercent float64,
) {

	totalRamStr := fmt.Sprintf("Total: %d GB", totalRam)
	usedRamStr := fmt.Sprintf("Used: %d GB, %f%c", usedRam, usedPercent, '%')
	availableRamStr := fmt.Sprintf("Available: %d GB", availableRam)

	ramListItems := []string{
		totalRamStr,
		usedRamStr,
		availableRamStr,
	}

	ramList = widget.NewList(
		func() int {
			return len(ramListItems)
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(index int, obj fyne.CanvasObject) {
			if label, ok := obj.(*widget.Label); ok {
				label.SetText(ramListItems[index])
			}
		},
	)

}
