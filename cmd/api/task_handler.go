package main

import (
	"errors"
	"net/http"

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

	err = app.stores.Tasks.Create(task)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// t, err := app.client.Task.
	// Create().
	// 	SetTitle(taskD.Title).
	// 	SetDescription(taskD.Description).
	// 	SetPriority(task.Priority(taskD.Priority)).
	// 	SetStatus(task.Status(taskD.Status)).
	// 	Save(context.Background())

	// if err != nil {
	// 	app.logger.Println(err)
	// 	app.serverErrorResponse(w, r, err)
	// 	// return nil, fmt.Errorf("failed creating task: %w", err)
	// }

	app.createdResponse(w, task)
}

func (app *application) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	// t, err := app.client.Task.Get(context.Background(), int(id))
	// if err != nil {
	// 	// var notFoundError *ent.NotFoundError
	// 	if _, ok := err.(*ent.NotFoundError); ok{
	// 		app.notFoundResponse(w, r)
	// 		return
	// 	}else{
	// 		app.serverErrorResponse(w, r, err)
	// 		return
	// 	}
	// }

	task, err := app.stores.Tasks.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.okResponse(w, task)
}

func (app *application) listTasksHandler(w http.ResponseWriter, r *http.Request) {

	// To keep things consistent with our other handlers, we'll define an input struct 
	// to hold the expected values from the request query string.
	var input struct {
		Title       string            
		Priority       string            
		Status       string            
		Description string   
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Priority = app.readString(qs, "priority", "")
	input.Status = app.readString(qs, "status", "")
	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 50, v)
	input.Sort = app.readString(qs, "sort", "id")

	input.SortSafelist = []string{
		"id", "-id", "title", "-title", "priority", "status"}


	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	tasks, err := app.stores.Tasks.GetAll(input.Title, input.Priority, input.Status, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	app.okResponse(w, tasks)
}

func (app *application) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	task, err := app.stores.Tasks.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title       *string            `json:"title"`
		Description *string            `json:"description"`
		Priority    *data.TaskPriority `json:"priority"`
		Status      *data.TaskStatus   `json:"status"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = *input.Description
	}
	if input.Priority != nil {
		task.Priority = *input.Priority
	}
	if input.Status != nil {
		task.Status = *input.Status
	}

	v := validator.New()

	if data.ValidateTask(v, task); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.stores.Tasks.Update(task)
	if err != nil {
		switch  {
		case errors.Is(err, data.ErrEditConflict):
			app.errorResponse(w, r, http.StatusConflict, 
				"unable to update the record due to an edit conflict, please try again")
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.okResponse(w, task)
}

func (app *application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.stores.Tasks.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.okResponse(w, "task successfully deleted")
}

// func (app *application) CreateTask() (*ent.Task, error) {
// 	ctx := context.Background()
// 	t, err := app.client.Task.
// 		Create().
// 		SetTitle("Task one").
// 		SetDescription("Task Description").
// 		SetPriority("low").
// 		SetStatus("todo").
// 		Save(ctx)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed creating task: %w", err)
// 	}

// 	app.logger.Println("task was created: ", t)
// 	return t, nil
// }
