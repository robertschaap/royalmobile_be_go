package router

import (
	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/controllers"
)

// CreateRouter initialises Mux Router and sets up routes
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", controllers.GetIndex)
	api.HandleFunc("/cart/{cartID}", controllers.GetCart)
	api.HandleFunc("/cart/{cartID}/item", controllers.AddCartItem)
	api.HandleFunc("/cart/{cartID}/item/{itemID}", controllers.DeleteCartItem)
	api.HandleFunc("/cart/order", controllers.PostOrder)
	api.HandleFunc("/product/{modelID}", controllers.GetProduct)
	api.HandleFunc("/products", controllers.GetProducts)
	api.HandleFunc("/subscriptions", controllers.GetSubscriptions)

	r.PathPrefix("/").HandlerFunc(controllers.GetError)

	return r
}
