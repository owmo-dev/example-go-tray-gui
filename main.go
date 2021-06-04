package main

import (
	"github.com/getlantern/systray"
	"github.com/granmo/ExampleTrayGUI/tray"
	"github.com/granmo/ExampleTrayGUI/views"
)

func main() {
	views := views.Get()
	defer views.WaitGroup.Wait()
	systray.Run(tray.OnReady, tray.OnQuit)
}
