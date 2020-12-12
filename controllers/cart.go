package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/models"
	"github.com/robertschaap/royalmobile_go_be/server"
)

type cart struct {
	ID     string     `json:"id"`
	Items  []cartItem `json:"items"`
	Totals cartTotals `json:"totals"`
}

type cartItem struct {
	product      models.Product
	subscription models.Subscription
	totals       cartTotals
}

type cartTotals struct {
	monthlyPrice string
	oneTimePrice string
}

// GetCart returns a shopping cart with products
func GetCart(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}

// AddCartItem returns a shopping cart with products
func AddCartItem(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}

// DeleteCartItem returns a shopping cart with products
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}

// PostOrder returns a shopping cart with products
func PostOrder(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(nil)
	res.JSON(w)
}
