package models

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
	monthlyPrice string
	oneTimePrice string
}
