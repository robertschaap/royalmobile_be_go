package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetCart returns a shopping cart with products
func GetCart(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}
