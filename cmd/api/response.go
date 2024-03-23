package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type envelope struct {
	Status string         `json:"status" xml:"status"`
	Code   int            `json:"code" xml:"code"`
	Data   map[string]any `json:"data" xml:"data"`
}

// WriteJSON Define a WriteJSON() helper for sending responses. This takes the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON, and a
// header map containing any additional HTTP headers we want to include in the response.
func (app *application) WriteJSON(w http.ResponseWriter, r *http.Request, status int, data any, headers http.Header) {
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Error("Failed to marshal data", slog.String("error", err.Error()))
		app.ResponseErrorInternalServerError(w, r, map[string]any{"error": "the server encountered a problem and could not process your request"}, nil)
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		app.logger.Error("Failed to write response", slog.String("error", err.Error()))
		app.ResponseErrorInternalServerError(w, r, map[string]any{"error": "the server encountered a problem and could not process your request"}, nil)
	}
}
