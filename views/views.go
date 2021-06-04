package views

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

const PORT = 8080
const HOST = "localhost"

var once sync.Once

//go:embed www
var fs embed.FS

type Views struct {
	list      map[string]*View
	WaitGroup *sync.WaitGroup
	Shutdown  chan bool
}

type View struct {
	url    string
	width  int
	height int
	isOpen bool
}

var views *Views

func Get() *Views {
	once.Do(func() {
		l := make(map[string]*View)

		l["Hello"] = &View{
			url:    fmt.Sprintf("http://%s/www/index.html", fmt.Sprintf("%s:%d", HOST, PORT)),
			width:  600,
			height: 280,
		}

		l["Google"] = &View{
			url:    "https://www.google.com/",
			width:  960,
			height: 800,
		}

		views = &Views{
			list:      l,
			WaitGroup: &sync.WaitGroup{},
			Shutdown:  make(chan bool),
		}

		views.WaitGroup.Add(1)
		go func(*Views) {
			defer views.WaitGroup.Done()
			ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", HOST, PORT))
			if err != nil {
				log.Fatal(err)
			}
			defer ln.Close()

			go func() {
				_ = http.Serve(ln, http.FileServer(http.FS(fs)))
			}()
			<-views.Shutdown
		}(views)
	})
	return views
}

func (v *Views) getView(name string) (*View, error) {
	view, ok := v.list[name]
	if !ok {
		return nil, fmt.Errorf("View '%s' not found", name)
	}
	if view.isOpen {
		return nil, fmt.Errorf("View is already open")
	}
	return view, nil
}
