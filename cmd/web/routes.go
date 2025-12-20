package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *Application) routes() http.Handler{
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{snippetId}",  dynamic.ThenFunc(app.viewSnippet))
	mux.Handle("GET /snippet/create",  dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create",  dynamic.ThenFunc(app.snippetCreatePost))

	middelwareChain := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return middelwareChain.Then(mux)
}