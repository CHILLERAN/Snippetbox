package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	//files... takes the values out of the slice and passes it on as individual arguments
	ts, err := template.ParseFiles(files...) 

	if err != nil {
		app.ServerError(w, r, err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.ServerError(w, r, err)
	}
}

func (app *Application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("snippetId"))

	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "Viewing snippet with id %d", id)
}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing create snippet form"))
}

func (app *Application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(201)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post created"))
}