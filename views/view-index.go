package views

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ctrlshiftmake/example-go-desktop-TrayGUI/config"
	"github.com/zserge/lorca"
)

type info struct {
	sync.Mutex
}

func (i *info) appVersion() string {
	i.Lock()
	defer i.Unlock()
	return config.ApplicationVersion
}

func (v *Views) OpenIndex() error {
	view, err := v.getView("Hello")
	if err != nil {
		return err
	}

	v.WaitGroup.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		ui, err := lorca.New("", "", view.width, view.height)
		if err != nil {
			log.Fatal(err)
		}
		defer ui.Close()

		i := info{}

		err = ui.Bind("appVersion", i.appVersion)
		if err != nil {
			log.Fatal(err)
		}

		err = ui.Load(view.url)
		if err != nil {
			log.Fatal(err)
		}

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
