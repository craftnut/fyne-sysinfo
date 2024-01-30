TITLE "Windows Package Script"
echo This assumes you have Fyne installed
TIMEOUT /T 3
call go get github.com/shirou/gopsutil/v3/
call go get github.com/wille/osutil
call go get fyne.io/fyne/v2
call fyne package -os windows -icon .\icon.png 