package main

import (
	"database/sql"
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
	DB          *sql.DB
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "patterns.db", "DSN")
	flag.Parse()

	// get database
	db, err := initDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	app.DB = db

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World!")
	// })

	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(), IdleTimeout: 30 * time.Second, ReadTimeout: 30 * time.Second, ReadHeaderTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second,
	}

	fmt.Println("Web application starting on port", port)

	// err := http.ListenAndServe(port, nil)
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
