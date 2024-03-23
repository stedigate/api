package api

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) statsHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"db": map[string]any{
			"Max Open Connections": app.models.DB.Stats().MaxOpenConnections,
			"Open Connections":     app.models.DB.Stats().OpenConnections,
			"Max Idle Closed":      app.models.DB.Stats().MaxIdleClosed,
			"Max Lifetime Closed":  app.models.DB.Stats().MaxLifetimeClosed,
			"Max Idle Time":        app.models.DB.Stats().MaxIdleTimeClosed,
			"Wait Count":           app.models.DB.Stats().WaitCount,
			"Wait Duration":        app.models.DB.Stats().WaitDuration.String(),
			"In Use":               app.models.DB.Stats().InUse,
			"Idle":                 app.models.DB.Stats().Idle,
		},
	}

	app.ResponseSuccessOk(w, r, data, nil)
}
