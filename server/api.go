package server

import (
	"encoding/json"
	"net/http"
)

// apiResponse is the actual base response we return to a client. This is private to prevent
// consumers from modifying any of the fields on the APIResponse struct without explicitly
// callig the provided methods.
type apiResponseJSON struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

// APIResponse defines basic response structure that is called with either Success or Error
// and sent to the client with the JSON method.
type APIResponse struct {
	status  string
	data    interface{}
	message interface{}
}

// Success response returns a successful response with the passed data
func (r *APIResponse) Success(data interface{}) *APIResponse {
	r.status = "success"
	r.data = data
	r.message = nil

	return r
}

// Error response returns an error response with the passed error message
func (r *APIResponse) Error(message string) *APIResponse {
	r.status = "error"
	r.message = message
	r.data = nil

	return r
}

// JSON converts the APIReponse struct to JSON and calls http.ResponseWriter with it
func (r *APIResponse) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	res := apiResponseJSON{}
	res.Status = r.status
	res.Data = r.data
	res.Message = r.message

	json.NewEncoder(w).Encode(res)
}
