package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// ProductVariant struct is a struct
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

// GetProducts reads products from disk or returns an empty slice
func GetProducts() []Product {
	products := []Product{}

	data, err := ioutil.ReadFile("./stubs/products.json")

	if err == nil {
		json.Unmarshal(data, &products)
	}

	return products
}

// GetProduct reads products from disk or returns an empty slice
func GetProduct(modelID string) (Product, error) {
	products := []Product{}

	data, err := ioutil.ReadFile("./stubs/products.json")

	if err == nil {
		json.Unmarshal(data, &products)
	}

	var product Product

	for _, v := range products {
		if v.ModelID == modelID {
			product = v
			return product, nil
		}
	}

	return product, errors.New("Could not get product")
}
