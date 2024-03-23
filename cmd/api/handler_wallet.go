package api

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) viewWalletHandler(w http.ResponseWriter, r *http.Request) {
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
func (app *application) createWalletHandler(w http.ResponseWriter, r *http.Request) {
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
func (app *application) updateWalletHandler(w http.ResponseWriter, r *http.Request) {
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
func (app *application) listWalletHandler(w http.ResponseWriter, r *http.Request) {
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
func (app *application) deleteWalletHandler(w http.ResponseWriter, r *http.Request) {
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
