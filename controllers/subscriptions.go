package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/server"
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

func getSubscriptionsStub() []Subscription {
	data, err := ioutil.ReadFile("./stubs/subscriptions.json")

	if err != nil {
		fmt.Print(err)
	}

	var subscriptions []Subscription

	err = json.Unmarshal(data, &subscriptions)

	if err != nil {
		fmt.Println("error:", err)
	}

	return subscriptions
}

// GetSubscriptions returns a list of all available subscriptions
func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{
		Status:  "success",
		Data:    getSubscriptionsStub(),
		Message: "",
	}

	res.JSON(w)
}
