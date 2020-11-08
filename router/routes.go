package router

import (
	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/controllers"
)

// CreateRouter initialises Mux Router and sets up routes
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	// TOOD: mux has a way to simplify the /api prefix, do it
	r.HandleFunc("/api/", controllers.GetIndex)
	r.HandleFunc("/api/cart/{cartId}", controllers.GetCart)
	r.HandleFunc("/api/cart/{cartId}/item", controllers.AddCartItem)
	r.HandleFunc("/api/cart/{cartId}/item/{itemId}", controllers.DeleteCartItem)
	r.HandleFunc("/api/cart/order", controllers.PostOrder)
	r.HandleFunc("/api/product/{id}", controllers.GetProduct)
	r.HandleFunc("/api/products", controllers.GetProducts)
	r.HandleFunc("/api/subscriptions", controllers.GetSubscriptions)

	return r
}
