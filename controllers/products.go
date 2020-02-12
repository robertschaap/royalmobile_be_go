package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetProducts returns a list of all available phones and basic specifications.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{
		Status:  "success",
		Data:    "",
		Message: "",
	}

	res.JSON(w)
}
