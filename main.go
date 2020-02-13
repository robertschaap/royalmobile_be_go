package main

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/router"
)

func main() {
	r := router.CreateRouter()
	http.ListenAndServe(":4000", r)
}
