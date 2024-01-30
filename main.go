package main

// main.go gets information into variables and creates the main menu window
import (
	"fmt"
	"math"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/wille/osutil"

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

	memUsedGb := int(
		math.Round(float64(mem.Used/1.074e+9) + .5),
	)

	cpuVendor := cpu[0].VendorID
	cpuModel := cpu[0].ModelName
	cpuCores := cpu[0].Cores
	cpuSpeed := cpu[0].Mhz

	osName := osutil.Name
	osVer := osutil.GetVersion()

	host, _ := host.Info()

	hostname := host.Hostname
	hostArch := host.KernelArch

	launchMainWindow(memFreeGb, memAvailableGb, memUsedGb, mem.UsedPercent, cpuVendor, cpuModel, cpuCores, cpuSpeed, osName, osVer, hostname, hostArch)
}

func launchMainWindow(totalMem int, availableMem int, usedMem int, usedMemPercent float64, cpuVendor string, cpuModel string, cpuCores int32, cpuSpeed float64, osName string, osVer string, hostname string, hostArch string) {

	fmt.Println(totalMem)
	fmt.Println(availableMem)

	app := app.New()
	mainWindow := app.NewWindow("SysInfo")

	cpuButton := widget.NewButton("CPU Info", func() {
		info.CpuInfo(cpuVendor, cpuModel, cpuCores, cpuSpeed, app)
	})
	memButton := widget.NewButton("RAM Info", func() {
		info.RamInfo(totalMem, availableMem, usedMem, usedMemPercent, app)
	})
	osButton := widget.NewButton("OS Info", func() {
		info.OsInfo(osName, osVer, hostname, hostArch, app)
	})

	layout := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 25)), cpuButton, memButton, osButton)

	mainWindow.SetContent(layout)
	mainWindow.Resize(fyne.NewSize(250, (25 * 3)))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
