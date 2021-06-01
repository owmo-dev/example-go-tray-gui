package views

import (
	"log"
	"sync"

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

		view.isOpen = true

		select {
		case <-ui.Done():
		case <-v.Shutdown:
		}

		view.isOpen = false

	}(v.WaitGroup)

	return nil
}
