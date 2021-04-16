package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/models"
	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetCart takes a UUIDv4 string "cartID" and returns a Cart or error
func GetCart(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["cartID"]

	cart, err := models.GetCartByID(id)

	res := server.APIResponse{}

	if err == nil {
		res.Success(cart)
	} else {
		res.Error("Cart could not be found")
	}

	res.JSON(w)
}

type addCartItemBody struct {
	VariantID      string `json:"variantId"`
	SubscriptionID string `json:"subscriptionId"`
}

// AddCartItem takes a UUIDv4 string "cartID" or the keyword "new" and returns a Cart or error
func AddCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["cartID"]

	// TODO also parse x-www-form-urlencoded instead of just json
	var body addCartItemBody
	err := json.NewDecoder(r.Body).Decode(&body)

	cart, err := models.AddCartItem(id, body.VariantID, body.SubscriptionID)

	res := server.APIResponse{}

	if err == nil {
		res.Success(cart)
	} else {
		res.Error("Could not add cart item")
	}

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
