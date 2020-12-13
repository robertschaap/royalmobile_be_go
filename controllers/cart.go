package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/models"
	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetCart takes a UUIDv4 string "cartID" and returns a Cart or error
func GetCart(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(models.GetCartByID("new-cart"))
	res.JSON(w)
}

// AddCartItem takes a UUIDv4 string "cartID" or the keyword "new" and returns a Cart or error
func AddCartItem(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}

// DeleteCartItem takes a UUIDv4 string "cartID" and "itemID" and returns a Cart without the item to delete or error
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}

// PostOrder takes a UUIDv4 string "cartID" and returns a Cart if succesful or an error
func PostOrder(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}
