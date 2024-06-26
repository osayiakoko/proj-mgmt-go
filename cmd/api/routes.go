package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Post("/v1/tasks", app.createTaskHandler)
	router.Get("/v1/tasks/{id}", app.showTaskHandler)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	return router
}
