package controllers

import (
	"net/http"
)

//BalanceController can only display the logged in user's balance
type Balance struct {
}

func (x *Balance) Get(c *gin.Context) {
	return http.StatusNotImplemented, nil
}
