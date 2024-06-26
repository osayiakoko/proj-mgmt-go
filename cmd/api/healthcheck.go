package main

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	app.okResponse(w, data)
}
