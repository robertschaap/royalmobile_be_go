package models

import "errors"

// Cart struct is a struct
type Cart struct {
	ID     string     `json:"id"`
	Items  []cartItem `json:"items"`
	Totals cartTotals `json:"totals"`
}

type cartItem struct {
	product      Product
	subscription Subscription
	totals       cartTotals
}

type cartTotals struct {
	MonthlyPrice string `json:"monthly_price"`
	OneTimePrice string `json:"one_time_price"`
}

var carts = []Cart{
	{
		ID:     "new-cart",
		Items:  make([]cartItem, 0),
		Totals: cartTotals{},
	},
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
