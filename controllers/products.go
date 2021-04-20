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
	res.Success(models.GetProducts()).JSON(w)
}

// GetProduct takes a modelID and returns a Product or error
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["modelID"]

	product, err := models.GetProduct(id)
	res := server.APIResponse{}

	if err == nil {
		res.Success(product).JSON(w)
	} else {
		res.Error("Could not get product").JSON(w)
	}
}
