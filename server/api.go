package server

import (
	"encoding/json"
	"net/http"
)

// APIResponse defines basic response structure.
// Status is either success or error.
// Data is the response data.
// Message is an optional error message.
type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// JSON converts the APIReponse struct to JSON and sends it as response.
func (r *APIResponse) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}
