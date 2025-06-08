package main

import (
	"context"
	"embed"
	"io/fs"
	"net/http"
	"sync/atomic"
	"templhtmxtests/views"
)

func main() {
	config := Config{}
	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.Index().Render(context.Background(), w)
	})
	server.HandleFunc("/count", config.countHandler)
	server.HandleFunc("/getcount", config.getCountHandler)

	// serve /static folder
	staticFS, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	http.ListenAndServe(":8080", server)
}

//go:embed static
var static embed.FS

type Config struct {
	Count atomic.Int32
}

func (c *Config) countHandler(w http.ResponseWriter, r *http.Request) {
	count := c.Count.Add(1)
	views.CounterResults(int(count)).Render(context.Background(), w)
}

func (c *Config) getCountHandler(w http.ResponseWriter, r *http.Request) {
	count := c.Count.Load()
	views.CounterResults(int(count)).Render(context.Background(), w)
}
