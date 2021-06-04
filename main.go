package main

import (
	"github.com/getlantern/systray"
	"github.com/granmo/example-go-desktop-TrayGUI/tray"
	"github.com/granmo/example-go-desktop-TrayGUI/views"
)

func main() {
	views := views.Get()
	defer views.WaitGroup.Wait()
	systray.Run(tray.OnReady, tray.OnQuit)
}
