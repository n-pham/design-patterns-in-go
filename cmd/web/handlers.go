package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World!")
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	if page != "favicon.ico" {
		app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
	}
}
