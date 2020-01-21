package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

//BalanceController can only display the logged in user's balance
type Balance struct {
}

func (x *Balance) Get(c *gin.Context) {
	return http.StatusNotImplemented, nil
}
