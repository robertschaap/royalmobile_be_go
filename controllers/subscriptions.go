package controllers

import (
	"net/http"

	"github.com/robertschaap/royalmobile_go_be/models"
	"github.com/robertschaap/royalmobile_go_be/server"
)

// GetSubscriptions returns a list of Subscriptions or an empty list
func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	res := server.APIResponse{}
	res.Success(models.GetSubscriptions()).JSON(w)
}
