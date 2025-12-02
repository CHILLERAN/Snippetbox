package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CHILLERAN/Snippetbox/internal/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
    
    snippets, err := app.snippets.Latest()

    if err != nil {
        app.ServerError(w, r, err)
        return
    }

	data := app.newTemplateData(r)
    data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *Application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("snippetId"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.ServerError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
    data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing create snippet form"))
}

func (app *Application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(201)

	title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
    expires := 7

	id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.ServerError(w, r, err)
        return
    }

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}