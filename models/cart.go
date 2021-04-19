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

func NewCart() Cart {
	return Cart{
		uuid.NewString(),
		make([]cartItem, 0),
		cartTotals{"0", "0"},
	}
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
	cart := NewCart()
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

	// TODO: figure out if this is idiomatic Go
	return Cart{}, errors.New("Cart could not be found")
}

// AddCartItem adds a cart item
func AddCartItem(cartID string, variantID string, subscriptionID string) (Cart, error) {
	usedCartID := cartID

	if cartID == "new" {
		usedCartID = createCart().ID
	}

	// TODO get or create new cart
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
