package tray

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/ctrlshiftmake/example-tray-gui/icon"
	"github.com/ctrlshiftmake/example-tray-gui/views"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func OnReady() {
	systray.SetIcon(icon.Data)

	mHelloWorld := systray.AddMenuItem("Hello, World!", "Opens a simple HTML Hello, World")
	systray.AddSeparator()
	mGoogleBrowser := systray.AddMenuItem("Google in Browser", "Opens Google in a normal browser")
	mGoogleEmbed := systray.AddMenuItem("Google in Window", "Opens Google in a custom window")
	systray.AddSeparator()
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
		case <-mGoogleBrowser.ClickedCh:
			err := open.Run("https://www.google.com")
			if err != nil {
				fmt.Println(err)
			}
		case <-mGoogleEmbed.ClickedCh:
			err := views.Get().OpenGoogle()
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
