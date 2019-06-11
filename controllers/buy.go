package controllers

import "github.com/louisevanderlith/mango/control"

type BuyController struct {
	control.APIController
}

func NewBuyCtrl(ctrlmap *control.ControllerMap) *BuyController {
	result := new(BuyController)
	result.SetInstanceMap(ctrlmap)

	return result
}

func (req *BuyController) Post() {

}
