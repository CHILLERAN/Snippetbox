package main

import "net/http"

func (app *Application) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri = r.RequestURI
	)

	app.logger.Error(err.Error(), "method", method, "URI", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
 }