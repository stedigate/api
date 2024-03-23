package api

import (
	"fmt"
	"net/http"
)

func (app *application) logError(err error) {
	app.logger.Error(err.Error())
}

func (app *application) WriteErrorJson(w http.ResponseWriter, r *http.Request, status int, data map[string]any, headers http.Header) {
	res := envelope{
		Status: "FAILED",
		Code:   status,
		Data:   data,
	}

	app.WriteJSON(w, r, status, res, headers)
}

func (app *application) ResponseErrorInternalServerError(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	data := map[string]any{
		"message": "the server encountered a problem and could not process your request",
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusInternalServerError, data, headers)
}

func (app *application) ResponseErrorNotFound(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	data := map[string]any{
		"message": "the requested resource could not be found",
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusNotFound, data, headers)
}

// ResponseErrorMethodNotAllowed method will be used to send a 405 Method Not Allowed
// status code and JSON response to the client.
func (app *application) ResponseErrorMethodNotAllowed(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusMethodNotAllowed, data, headers)
}

func (app *application) ResponseErrorBadRequest(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	data := map[string]any{
		"message": "the request used wrong method or contains malformed and/or illegal data",
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusBadRequest, data, headers)
}

func (app *application) ResponseErrorFailedValidation(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	data := map[string]any{
		"message": "validation failed",
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusUnprocessableEntity, data, headers)
}

// ResponseErrorEditConflict sends a JSON-formatted error message to the client with a 409 Conflict
// status code.
func (app *application) ResponseErrorEditConflict(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	message := "unable to update the record due to an edit conflict, please try again"
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusConflict, data, headers)
}

func (app *application) ResponseErrorRateLimitExceeded(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	message := "rate limited exceeded"
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusTooManyRequests, data, headers)
}

func (app *application) ResponseErrorInvalidCredentials(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	message := "invalid authentication credentials"
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusUnauthorized, data, headers)
}

func (app *application) ResponseErrorInvalidAuthenticationToken(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	w.Header().Set("WWWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusUnauthorized, data, headers)
}

func (app *application) ResponseErrorAuthenticationRequired(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	message := "you must be authenticated to access this resource"
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusUnauthorized, data, headers)
}

func (app *application) ResponseErrorInactiveAccount(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	message := "your user account must be activated to access this resource"
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusForbidden, data, headers)
}

func (app *application) ResponseErrorNotPermitted(w http.ResponseWriter, r *http.Request, errors map[string]any, headers http.Header) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	data := map[string]any{
		"message": message,
		"errors":  errors,
	}
	app.WriteErrorJson(w, r, http.StatusForbidden, data, headers)
}
