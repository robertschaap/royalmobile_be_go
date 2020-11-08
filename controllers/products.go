package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertschaap/royalmobile_go_be/server"
)

// ProductVariant struct denotes a variation of the device with different color and/or capacity
type ProductVariant struct {
	ID              int32  `json:"id"`
	VariantID       string `json:"variantId"`
	Color           string `json:"color"`
	ColorHex        string `json:"colorHex"`
	Capacity        string `json:"capacity"`
	IsInStock       bool   `json:"is_in_stock"`
	IsPreorder      bool   `json:"is_preorder"`
	RegularPrice    string `json:"regular_price"`
	DiscountedPrice string `json:"discounted_price"`
	HasDiscounts    bool   `json:"has_discounts"`
}

// Product struct denotes the base information of the device
type Product struct {
	ID           int32            `json:"id"`
	Manufacturer string           `json:"manufacturer"`
	Model        string           `json:"model"`
	ModelID      string           `json:"modelId"`
	Variants     []ProductVariant `json:"variants"`
}

func getProductsStubs() []Product {
	data, err := ioutil.ReadFile("./stubs/products.json")

	if err != nil {
		fmt.Print(err)
	}

	var products []Product

	err = json.Unmarshal(data, &products)

	if err != nil {
		fmt.Println("error:", err)
	}

	return products
}

// GetProducts returns a list of all available phones and basic specifications.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(getProductsStubs())
	res.JSON(w)
}

// GetProduct returns a single product
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var product Product

	for _, v := range getProductsStubs() {
		if v.ModelID == id {
			product = v
			break
		}
	}

	res := server.APIResponse{}
	res.Success(product)
	res.JSON(w)
}
