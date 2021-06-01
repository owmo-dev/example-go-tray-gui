package main

import (
	"grantmoore3d/example-go-desktop-TrayGUI/tray"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(tray.OnReady, tray.OnQuit)
}
