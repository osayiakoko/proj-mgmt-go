package main

import (
	"net/http"
)

func (app *application) logError(_ *http.Request, err error) {
	app.logger.Println(err)
}
