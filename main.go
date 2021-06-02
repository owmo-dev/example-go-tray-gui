package main

import (
	"grantmoore3d/example-go-desktop-TrayGUI/tray"
	"grantmoore3d/example-go-desktop-TrayGUI/views"

	"github.com/getlantern/systray"
)

func main() {
	defer views.Get().WaitGroup.Wait()
	systray.Run(tray.OnReady, tray.OnQuit)
}
