package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (r *ApiResponse) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	res := ApiResponse{"", "", ""}
	res.JSON(w)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	res := ApiResponse{"", "", ""}
	res.JSON(w)
}

func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	res := ApiResponse{"", "", ""}
	res.JSON(w)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", GetIndex)
	r.HandleFunc("/api/products", GetProducts)
	r.HandleFunc("/api/subscriptions", GetSubscriptions)
	http.ListenAndServe(":4000", r)
}
