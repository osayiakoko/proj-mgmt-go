package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Logger)
	// router.Use(middleware.Recoverer)
	router.Use(app.recoverer)
	router.Use(app.clientRateLimit)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	// TASKS route
	router.Post("/v1/tasks", app.createTaskHandler)
	router.Get("/v1/tasks", app.listTasksHandler)
	router.Get("/v1/tasks/{id}", app.getTaskHandler)
	router.Patch("/v1/tasks/{id}", app.updateTaskHandler)
	router.Delete("/v1/tasks/{id}", app.deleteTaskHandler)

	// USERS route
	router.Post("/v1/users", app.registerUserHandler)
	router.Put("/v1/users/activate", app.activateUserHandler)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	return router
}
