package api

import "net/http"

// ResponseSuccessOk Define a WriteJSON() helper for sending responses. This takes the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON, and a
// header map containing any additional HTTP headers we want to include in the response.
func (app *application) ResponseSuccessOk(w http.ResponseWriter, r *http.Request, data map[string]any, headers http.Header) {
	res := envelope{
		Status: "SUCCESS",
		Code:   200,
		Data:   data,
	}
	app.WriteJSON(w, r, http.StatusOK, res, headers)
}

// ResponseSuccessCreated Define a WriteJSON() helper for sending responses. This takes the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON, and a
// header map containing any additional HTTP headers we want to include in the response.
func (app *application) ResponseSuccessCreated(w http.ResponseWriter, r *http.Request, data map[string]any, headers http.Header) {
	res := envelope{
		Status: "OK",
		Code:   201,
		Data:   data,
	}
	app.WriteJSON(w, r, http.StatusCreated, res, headers)
}

// ResponseSuccessEmpty Define a WriteJSON() helper for sending responses. This takes the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON, and a
// header map containing any additional HTTP headers we want to include in the response.
func (app *application) ResponseSuccessEmpty(w http.ResponseWriter, r *http.Request, data map[string]any, headers http.Header) {
	app.WriteJSON(w, r, http.StatusNoContent, data, headers)
}
