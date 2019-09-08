package routers

import (
	"github.com/louisevanderlith/xchange/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	buyCtrl := &controllers.Buy{}
	balCtrl := &controllers.Balance{}
	e.JoinBundle("/", roletype.User, mix.JSON, buyCtrl, balCtrl)
}
