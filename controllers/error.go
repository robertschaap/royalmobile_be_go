package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetError handles all undefined routes and returns a generic response
func GetError(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Error("This API route does not exist")
	res.JSON(w)
}
