package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", controllers.GetIndex)
	r.HandleFunc("/api/products", controllers.GetProducts)
	r.HandleFunc("/api/subscriptions", controllers.GetSubscriptions)
	http.ListenAndServe(":4000", r)
}
