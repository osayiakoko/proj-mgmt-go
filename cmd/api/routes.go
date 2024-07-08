package main

import (
	"expvar"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/osayiakoko/project-mgmt-sys/internal/data"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	// router.Use(middleware.Recoverer)
	router.Use(app.recoverer)
	router.Use(app.metrics)
	router.Use(middleware.Logger)
	router.Use(middleware.RedirectSlashes)
	router.Use(app.enableCORS)
	router.Use(app.clientRateLimit)
	router.Use(app.authenticate)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	// TASKS route
	router.Post("/v1/tasks", app.requirePermission(data.TaskWritePermission, app.createTaskHandler))
	router.Get("/v1/tasks", app.requirePermission(data.TaskReadPermission, app.listTasksHandler))
	router.Get("/v1/tasks/{id}", app.requirePermission(data.TaskReadPermission, app.getTaskHandler))
	router.Patch("/v1/tasks/{id}", app.requirePermission(data.TaskWritePermission, app.updateTaskHandler))
	router.Delete("/v1/tasks/{id}", app.requirePermission(data.TaskWritePermission, app.deleteTaskHandler))

	// USERS route
	router.Post("/v1/users", app.registerUserHandler)
	router.Put("/v1/users/activate", app.activateUserHandler)

	// AUTHS route
	router.Post("/v1/auths/login", app.createAuthenticationTokenHandler)

	// Register a new GET /debug/vars endpoint pointing to the expvar handler.
	router.Handle("/debug/vars", expvar.Handler())

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	return router
}
