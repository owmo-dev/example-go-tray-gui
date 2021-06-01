package main

import (
	"github.com/ctrlshiftmake/example-tray-gui/tray"
	"github.com/ctrlshiftmake/example-tray-gui/views"
	"github.com/getlantern/systray"
)

func main() {
	views := views.Get()
	defer views.WaitGroup.Wait()
	systray.Run(tray.OnReady, tray.OnQuit)
}
