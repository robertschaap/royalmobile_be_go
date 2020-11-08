package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetIndex returns basic information about the API when called
func GetIndex(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success("Hello WOrld!")

	res.JSON(w)
}
