package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/osayiakoko/project-mgmt-sys/internal/data"
	"github.com/osayiakoko/project-mgmt-sys/internal/validator"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string            `json:"title"`
		Description string            `json:"description"`
		Priority    data.TaskPriority `json:"priority"`
		Status      data.TaskStatus   `json:"status"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the input struct to a new Movie struct.
	task := &data.Task{
		Title:       input.Title,
		Description: input.Description,
		Priority:    input.Priority,
		Status:      input.Status,
	}

	v := validator.New()

	if data.ValidateTask(v, task); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	app.createdResponse(w, input)

	// fmt.Fprintf(w, "create a new task")
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	task := data.Task{
		ID:          id,
		Title:       fmt.Sprintf("Task %d", id),
		Description: fmt.Sprintf("Description for task %d", id),
		Priority:    "high",
		Status:      "ongoing",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	app.okResponse(w, task)
}
