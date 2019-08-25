package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

//BalanceController can only display the logged in user's balance
type BalanceController struct {
}

func (req *BalanceController) Get(ctx context.Contexer) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
