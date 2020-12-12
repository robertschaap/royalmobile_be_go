package models

import (
	"encoding/json"
	"io/ioutil"
)

// Subscription struct denotes the base information of a subscription
type Subscription struct {
	ID             int32    `json:"id"`
	SubscriptionID string   `json:"subscriptionId"`
	DurationID     string   `json:"durationId"`
	Data           string   `json:"data"`
	BenefitsLong   []string `json:"benefits_long"`
	BenefitsShort  string   `json:"benefits_short"`
	RegularPrice   string   `json:"regular_price"`
}

// GetSubscriptions reads subscriptions from disk or returns an empty slice
func GetSubscriptions() []Subscription {
	subscriptions := []Subscription{}

	data, err := ioutil.ReadFile("./stubs/subscriptions.json")

	if err != nil {
		return subscriptions
	}

	err = json.Unmarshal(data, &subscriptions)

	if err != nil {
		return subscriptions
	}

	return subscriptions

}
