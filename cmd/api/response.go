package main

import (
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

const failed, success = "failed", "success"

func (app *application) response(w http.ResponseWriter, r *http.Request, status int, data any, resStatus string, message any, headers http.Header) {

	resBody := ApiResponse{
		Status:  resStatus,
		Message: message,
		Data:    data,
	}

	err := app.writeJSON(w, status, resBody, headers)

	if err != nil {
		// app.logger.Println(err)
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		// http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

// SUCCESS/PASS RESPONSES
func (app *application) successResponse(w http.ResponseWriter, status int, data any) {
	app.response(w, nil, status, data, success, nil, nil)
}

func (app *application) createdResponse(w http.ResponseWriter, data any) {
	app.successResponse(w, http.StatusCreated, data)
}

func (app *application) okResponse(w http.ResponseWriter, data any) {
	app.successResponse(w, http.StatusOK, data)
}

// func (app *application) failedResponse(w http.ResponseWriter, status int, message any) {
// 	app.response(w, nil, status, nil, failed, message, nil)
// }

// ERROR/FAIL RESPONSES
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	app.response(w, r, status, nil, failed, message, nil)
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(w, r, http.StatusConflict, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}
