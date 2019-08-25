package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type BuyController struct {
}

func (req *BuyController) Post(ctx context.Contexer) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
