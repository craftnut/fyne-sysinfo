package main

import (
	"fmt"
	"math"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	info "github.com/craftnut/fyne-sysinfo/info"
)

func main() {
	mem, _ := mem.VirtualMemory()
	cpu, _ := cpu.Info()

	memFreeGb := int(
		math.Round(float64(mem.Total/1.074e+9) + .5),
	)

	memAvailableGb := int(
		math.Round(float64(mem.Available/1.074e+9) + .5),
	)

	cpuVendor := cpu[0].VendorID
	cpuModel := cpu[0].ModelName

	launchMainWindow(memFreeGb, memAvailableGb, cpuVendor, cpuModel)
}

func launchMainWindow(totalMem int, availableMem int, cpuVendor string, cpuModel string) {

	fmt.Println(totalMem)
	fmt.Println(availableMem)

	app := app.New()
	mainWindow := app.NewWindow("SysInfo")

	cpuButton := widget.NewButton("CPU Info", func() { info.CpuInfo(cpuVendor, cpuModel, app) })
	memButton := widget.NewButton("RAM Info", func() {})
	osButton := widget.NewButton("OS Info", func() {})

	layout := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 50)), cpuButton, memButton, osButton)

	mainWindow.SetContent(layout)
	mainWindow.Resize(fyne.NewSize(250, 150))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
