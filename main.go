package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/server"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{
		Status:  "",
		Data:    "",
		Message: "",
	}
	res.JSON(w)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{
		Status:  "",
		Data:    "",
		Message: "",
	}

	res.JSON(w)
}

func getSubscriptions(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{
		Status:  "",
		Data:    "",
		Message: "",
	}

	res.JSON(w)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", getIndex)
	r.HandleFunc("/api/products", getProducts)
	r.HandleFunc("/api/subscriptions", getSubscriptions)
	http.ListenAndServe(":4000", r)
}
