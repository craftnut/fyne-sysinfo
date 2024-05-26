package info

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func OsInfo(
	osName string, osVer string,
	hostname string, hostArch string,
	platform string, app fyne.App,
) {

	osWindow := app.NewWindow("OS Info")

	osVer = fmt.Sprintf("version: %s", osVer)
	hostname = fmt.Sprintf("hostname: %s", hostname)

	osListItems := []string{
		osName, osVer,
		hostname, hostArch,
		platform,
	}

	osList := widget.NewList(
		func() int {
			return len(osListItems)
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},

		func(index int, obj fyne.CanvasObject) {
			if label, ok := obj.(*widget.Label); ok {
				label.SetText(osListItems[index])
			}
		},
	)

	layout := container.New(layout.NewGridLayout(1),
		osList,
	)

	osWindow.SetContent(layout)
	osWindow.Resize(fyne.NewSize(275, 225))
	osWindow.CenterOnScreen()
	osWindow.Show()

}
