package router

import (
	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/controllers"
)

// CreateRouter initialises Mux Router and sets up routes
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/", controllers.GetIndex)
	r.HandleFunc("/api/products", controllers.GetProducts)
	r.HandleFunc("/api/subscriptions", controllers.GetSubscriptions)

	return r
}
