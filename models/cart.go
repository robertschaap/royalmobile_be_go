package models

import (
	"errors"
	"strconv"

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

func updateCartTotals(cartItems []cartItem) cartTotals {
	monthlyPrice := 0
	oneTimePrice := 0

	for _, item := range cartItems {
		if p, err := strconv.Atoi(item.Totals.MonthlyPrice); err == nil {
			monthlyPrice += p
		}

		if p, err := strconv.Atoi(item.Totals.OneTimePrice); err == nil {
			oneTimePrice += p
		}
	}

	return cartTotals{
		strconv.Itoa(monthlyPrice),
		strconv.Itoa(oneTimePrice),
	}
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

	var item cartItem
	item.ID = uuid.NewString()

	if product, err := getProductByVariantID(variantID); err == nil {
		item.Product = product
	} else {
		return Cart{}, errors.New("Could not add cart item")
	}

	if subscription, err := getSubscriptionByID(subscriptionID); err == nil {
		item.Subscription = subscription
	} else {
		return Cart{}, errors.New("Could not add cart item")
	}

	item.Totals.MonthlyPrice = item.Subscription.RegularPrice
	item.Totals.OneTimePrice = item.Product.Variants[0].RegularPrice

	cart.Items = append(cart.Items, item)
	cart.Totals = updateCartTotals(cart.Items)

	for i, c := range carts {
		if c.ID == cart.ID {
			carts[i] = cart
			break
		}
	}

	return cart, nil
}

// DeleteCartItem deletes a cart item
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
	cart.Totals = updateCartTotals(cart.Items)

	for i, c := range carts {
		if c.ID == cart.ID {
			carts[i] = cart
			break
		}
	}

	return cart, nil
}

// PostOrder posts an order
func PostOrder(cartID string) (Cart, error) {
	if cart, err := GetCartByID(cartID); err == nil {
		return cart, nil
	} else {
		return Cart{}, errors.New("Could not post order")
	}
}
