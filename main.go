package main

// main.go gets information into variables and creates the main menu window
import (
	"fmt"
	"log"
	"math"
	"os/exec"

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

const byteDivider = 1074000000 // when divided by this, bytes become gibibytes

func main() {

	var osVer string
	var cpuThreads string

	mem, _ := mem.VirtualMemory()

	memFreeGb := int(
		math.Round(float64(mem.Total/byteDivider) + .5),
	)

	memAvailableGb := int(
		math.Round(float64(mem.Available/byteDivider) + .5),
	)

	memUsedGb := int(
		math.Round(float64(mem.Used/byteDivider) + .5),
	)

	cpu, _ := cpu.Info()
	fmt.Println(cpu)

	cpuVendor := cpu[0].VendorID
	cpuModel := cpu[0].ModelName

	cpuSpeed := cpu[0].Mhz

	host, _ := host.Info()

	hostname := host.Hostname
	hostArch := host.KernelArch
	platform := host.Platform

	osName := osutil.Name

	switch osName {
	case "Linux", "FreeBSD":
		osVer = host.KernelVersion
		nproc, err := exec.Command("nproc").Output()
		if err != nil {
			cpuThreads = "Unknown"
			log.Fatal(err)
		}
		cpuThreads = string(nproc)
	case "Windows", "macOS":
		osVer = osutil.GetVersion()
		cpuThreads = fmt.Sprintf("%d", cpu[0].Cores)
	default:
		osVer = "Unknown"
		cpuThreads = "Unknown"
	}

	launchMainWindow(
		memFreeGb, memAvailableGb, memUsedGb, mem.UsedPercent,
		cpuVendor, cpuModel, cpuThreads, cpuSpeed,
		osName, osVer, hostname, hostArch, platform,
	)

}

func launchMainWindow(
	totalMem int, availableMem int, usedMem int, usedMemPercent float64,
	cpuVendor string, cpuModel string, cpuThreads string, cpuSpeed float64,
	osName string, osVer string, hostname string, hostArch string, platform string,
) {

	fmt.Println(totalMem)
	fmt.Println(availableMem)

	app := app.New()
	mainWindow := app.NewWindow("SysInfo")

	cpuButton := widget.NewButton("CPU Info", func() {
		info.CpuInfo(cpuVendor, cpuModel, cpuThreads, cpuSpeed, app)
	})
	memButton := widget.NewButton("RAM Info", func() {
		info.RamInfo(totalMem, availableMem, usedMem, usedMemPercent, app)
	})
	osButton := widget.NewButton("OS Info", func() {
		info.OsInfo(osName, osVer, hostname, hostArch, platform, app)
	})

	layout := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 25)), cpuButton, memButton, osButton)

	mainWindow.SetContent(layout)
	mainWindow.Resize(fyne.NewSize(250, (25 * 3)))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()

}
