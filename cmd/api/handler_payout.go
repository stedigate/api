package api

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) viewPayoutHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.ResponseErrorNotFound(w, r, nil, nil)
		return
	}

	data := map[string]interface{}{
		"id": id,
	}

	app.ResponseSuccessOk(w, r, data, nil)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) createPayoutHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.ResponseErrorNotFound(w, r, nil, nil)
		return
	}

	data := map[string]interface{}{
		"id": id,
	}

	app.ResponseSuccessOk(w, r, data, nil)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) listPayoutHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	app.ResponseSuccessCreated(w, r, data, nil)
}
