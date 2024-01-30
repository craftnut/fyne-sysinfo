package info

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func OsInfo(osName string, osVer string, hostname string, hostArch string, app fyne.App) {

	osWindow := app.NewWindow("OS Info")

	nameLabel := widget.NewLabel(osName)
	verLabel := widget.NewLabel("Version: " + osVer)
	hostLabel := widget.NewLabel(hostname)
	archLabel := widget.NewLabel(hostArch)

	layout := container.New(layout.NewGridLayout(1),
		hostLabel,
		nameLabel,
		verLabel,
		archLabel,
	)

	osWindow.SetContent(layout)
	osWindow.CenterOnScreen()
	osWindow.Show()

}
