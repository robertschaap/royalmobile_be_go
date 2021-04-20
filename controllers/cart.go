package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

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
		res.Success(cart).JSON(w)
	} else {
		res.Error("Cart could not be found").JSON(w)
	}
}

type addCartItemBody struct {
	VariantID      string `json:"variantId"`
	SubscriptionID string `json:"subscriptionId"`
}

func decodeRequestBody(r *http.Request, target interface{}) error {
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		return json.NewDecoder(r.Body).Decode(target)
	}

	if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		rt := reflect.TypeOf(target).Elem()
		rv := reflect.ValueOf(target).Elem()

		if rv.Kind().String() != "struct" {
			return errors.New("Could not decode request body")
		}

		// Loop through fields in the target struct and read values from the request
		// into it based on the field tag from the target struct
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			f := rv.FieldByName(field.Name)
			ptr := f.Addr().Interface().(*string)

			*ptr = r.FormValue(field.Tag.Get("json"))
		}

		return nil
	}

	return errors.New("Could not decode request body")
}

// AddCartItem takes a UUIDv4 string "cartID" or the keyword "new" and returns a Cart or error
func AddCartItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["cartID"]
	res := server.APIResponse{}

	var body addCartItemBody
	err := decodeRequestBody(r, &body)

	if err != nil {
		res.Error("Could not add cart item").JSON(w)
		return
	}

	cart, err := models.AddCartItem(id, body.VariantID, body.SubscriptionID)

	if err != nil {
		res.Error("Could not add cart item").JSON(w)
		return
	}

	res.Success(cart).JSON(w)
}

// DeleteCartItem takes a UUIDv4 string "cartID" and "itemID" and returns a Cart without the item to delete or error
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartID := mux.Vars(r)["cartID"]
	itemID := mux.Vars(r)["itemID"]
	res := server.APIResponse{}

	cart, err := models.DeleteCartItem(cartID, itemID)

	if err != nil {
		res.Error("Could not delete cart item").JSON(w)
	}

	res.Success(cart).JSON(w)
}

// PostOrder takes a UUIDv4 string "cartID" and returns a Cart if succesful or an error
func PostOrder(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil).JSON(w)
}
