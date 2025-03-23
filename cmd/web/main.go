package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.Parse()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World!")
	// })

	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(), IdleTimeout: 30 * time.Second, ReadTimeout: 30 * time.Second, ReadHeaderTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second,
	}

	fmt.Println("Web application starting on port", port)

	// err := http.ListenAndServe(port, nil)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
