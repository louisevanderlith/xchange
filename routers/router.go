package routers

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/xchange/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilters(s, host)

	beego.Router("/v1/buy", controllers.NewBuyCtrl(ctrlmap), "post:Post")
	beego.Router("/v1/balance", controllers.NewBalanceCtrl(ctrlmap), "get:Get")
}

func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.User
	emptyMap["GET"] = roletype.User

	ctrlmap.Add("/v1/buy", emptyMap)
	ctrlmap.Add("/v1/balance", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
