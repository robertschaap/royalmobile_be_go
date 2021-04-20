package models

import (
	"errors"

	"github.com/google/uuid"
)

// Cart struct is a struct
type Cart struct {
	ID     string     `json:"id"`
	Items  []cartItem `json:"items"`
	Totals cartTotals `json:"totals"`
}

type cartItem struct {
	ID           string       `json:"id"`
	Product      Product      `json:"product"`
	Subscription Subscription `json:"subscription"`
	Totals       cartTotals   `json:"totals"`
}

type cartTotals struct {
	MonthlyPrice string `json:"monthly_price"`
	OneTimePrice string `json:"onetime_price"`
}

var carts = []Cart{
	{
		ID:     "new-cart",
		Items:  make([]cartItem, 0),
		Totals: cartTotals{"0", "0"},
	},
}

func createCart() Cart {
	cart := Cart{
		uuid.NewString(),
		make([]cartItem, 0),
		cartTotals{"0", "0"},
	}

	carts = append(carts, cart)

	return cart
}

// GetCartByID gets a cart by ID
func GetCartByID(cartID string) (Cart, error) {
	for _, v := range carts {
		if v.ID == cartID {
			return v, nil
		}
	}

	return Cart{}, errors.New("Cart could not be found")
}

// AddCartItem adds a cart item
func AddCartItem(cartID string, variantID string, subscriptionID string) (Cart, error) {
	usedCartID := cartID

	if cartID == "new" {
		usedCartID = createCart().ID
	}

	cart, err := GetCartByID(usedCartID)

	if err != nil {
		return Cart{}, errors.New("Could not add cart item")
	}

	product, err := getProductByVariantID(variantID)

	if err != nil {
		return Cart{}, errors.New("Could not add cart item")
	}

	subscription, err := getSubscriptionByID(subscriptionID)

	if err != nil {
		return Cart{}, errors.New("Could not add cart item")
	}

	item := cartItem{
		uuid.NewString(),
		product,
		subscription,
		cartTotals{
			subscription.RegularPrice,
			product.Variants[0].RegularPrice,
		},
	}

	cart.Items = append(cart.Items, item)

	var cartTotals cartTotals

	// Wrong values because of string concatenation
	for _, item := range cart.Items {
		cartTotals.MonthlyPrice += item.Totals.MonthlyPrice
		cartTotals.OneTimePrice += item.Totals.OneTimePrice
	}

	cart.Totals = cartTotals

	for i, c := range carts {
		if c.ID == cart.ID {
			carts[i] = cart
			break
		}
	}

	return cart, nil
}

func DeleteCartItem(cartID string, itemID string) (Cart, error) {
	cart, err := GetCartByID(cartID)

	if err != nil {
		return Cart{}, errors.New("Could not delete cart item")
	}

	index := -1
	for i, item := range cart.Items {
		if item.ID == itemID {
			index = i
			break
		}
	}

	if index == -1 {
		return Cart{}, errors.New("Could not delete cart item")
	}

	cart.Items = append(cart.Items[:index], cart.Items[index+1:]...)

	var cartTotals cartTotals

	// Wrong values because of string concatenation
	for _, item := range cart.Items {
		cartTotals.MonthlyPrice += item.Totals.MonthlyPrice
		cartTotals.OneTimePrice += item.Totals.OneTimePrice
	}

	cart.Totals = cartTotals

	for i, c := range carts {
		if c.ID == cart.ID {
			carts[i] = cart
			break
		}
	}

	return cart, nil
}
