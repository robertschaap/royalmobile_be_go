package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetSubscriptions returns a list of all available subscriptions
func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{
		Status:  "success",
		Data:    "",
		Message: "",
	}

	res.JSON(w)
}
