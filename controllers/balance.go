package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

//BalanceController can only display the logged in user's balance
type BalanceController struct {
	xontrols.APICtrl
}

func (req *BalanceController) Get() {

}
