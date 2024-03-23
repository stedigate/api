package api

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) viewInvoiceHandler(w http.ResponseWriter, r *http.Request) {
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
func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	app.ResponseSuccessCreated(w, r, data, nil)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) listInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	app.ResponseSuccessCreated(w, r, data, nil)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) deleteInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	app.ResponseSuccessCreated(w, r, data, nil)
}
