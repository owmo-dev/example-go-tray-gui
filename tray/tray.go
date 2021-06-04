package tray

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/getlantern/systray"
	"github.com/granmo/ExampleTrayGUI/icon"
	"github.com/granmo/ExampleTrayGUI/views"
)

func OnReady() {
	systray.SetIcon(icon.Data)

	mHelloWorld := systray.AddMenuItem("Hello, World!", "Opens a simple HTML Hello, World")

	mQuit := systray.AddMenuItem("Quit", "Quit example tray application")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-mHelloWorld.ClickedCh:
			err := views.Get().OpenIndex()
			if err != nil {
				fmt.Println(err)
			}
		case <-mQuit.ClickedCh:
			systray.Quit()
		case <-sigc:
			systray.Quit()
		}
	}
}

func OnQuit() {
	close(views.Get().Shutdown)
}
