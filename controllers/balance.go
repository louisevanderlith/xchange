package controllers

import "github.com/louisevanderlith/mango/control"

//BalanceController can only display the logged in user's balance
type BalanceController struct {
	control.APIController
}

func NewBalanceCtrl(ctrlmap *control.ControllerMap) *BalanceController {
	result := new(BalanceController)
	result.SetInstanceMap(ctrlmap)

	return result
}

func (req *BalanceController) Get() {
	
}