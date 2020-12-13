package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/models"
	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetProducts returns a list of Product or an empty list
func GetProducts(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(models.GetProducts())
	res.JSON(w)
}

// GetProduct takes a modelID and returns a Product or error
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["modelID"]

	var product models.Product

	for _, v := range models.GetProducts() {
		if v.ModelID == id {
			product = v
			break
		}
	}

	res := server.APIResponse{}
	res.Success(product)
	res.JSON(w)
}
