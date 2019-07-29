package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/xchange/controllers"

	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(poxy *droxolite.Epoxy) {
	//Buy
	buyCtrl := &controllers.BuyController{}
	buyGroup := droxolite.NewRouteGroup("buy", buyCtrl)
	buyGroup.AddRoute("/", "POST", roletype.User, buyCtrl.Post)
	poxy.AddGroup(buyGroup)

	//Balance
	balCtrl := &controllers.BalanceController{}
	balGroup := droxolite.NewRouteGroup("balance", balCtrl)
	balGroup.AddRoute("/", "GET", roletype.User, balCtrl.Get)
	poxy.AddGroup(balGroup)
	/*ctrlmap := EnableFilters(s, host)

	beego.Router("/v1/buy", controllers.NewBuyCtrl(ctrlmap), "post:Post")
	beego.Router("/v1/balance", controllers.NewBalanceCtrl(ctrlmap), "get:Get")*/
}

/*
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
*/
