package views

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/zserge/lorca"
)

func (v *Views) OpenGoogle() error {
	view, err := v.getView("Google")
	if err != nil {
		return err
	}

	v.WaitGroup.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		ui, err := lorca.New(view.url, "", view.width, view.height)
		if err != nil {
			log.Fatal(err)
		}
		defer ui.Close()

		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

		view.isOpen = true

		select {
		case <-sigc:
			v.Shutdown <- true
		case <-ui.Done():
		case <-v.Shutdown:
		}

		view.isOpen = false

	}(v.WaitGroup)

	return nil
}
