package main

import (
	"fmt"
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/router"
)

func main() {
	fmt.Print("starting server")
	r := router.CreateRouter()
	http.ListenAndServe(":4000", r)
}
