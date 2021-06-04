package main

import (
	"github.com/ctrlshiftmake/example-go-desktop-TrayGUI/tray"
	"github.com/ctrlshiftmake/example-go-desktop-TrayGUI/views"
	"github.com/getlantern/systray"
)

func main() {
	views := views.Get()
	defer views.WaitGroup.Wait()
	systray.Run(tray.OnReady, tray.OnQuit)
}
