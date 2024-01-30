TITLE "Package Script"
echo This assumes you have Fyne installed
echo Starts in 3 seconds
TIMEOUT /T 3
call go get github.com/shirou/gopsutil/v3/
call go get fyne.io/fyne/v2
call fyne package -os windows -icon .\icon.png 
call fyne package -os linux -icon .\icon.png 