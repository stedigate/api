package api

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"info": map[string]any{
			"status":      "available",
			"environment": app.config.App.Env,
			"version":     app.config.App.Version,
			"port":        app.config.App.Port,
		},
	}

	app.ResponseSuccessOk(w, r, data, nil)
}
