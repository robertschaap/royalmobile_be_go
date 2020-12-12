package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/robertschaap/royalmobile_go_be/model"
	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetProducts returns a list of all available phones and basic specifications.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(m.GetProducts())
	res.JSON(w)
}

// GetProduct returns a single product
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var product m.Product

	for _, v := range m.GetProducts() {
		if v.ModelID == id {
			product = v
			break
		}
	}

	res := server.APIResponse{}
	res.Success(product)
	res.JSON(w)
}
