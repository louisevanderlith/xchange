package routers

import (
	"github.com/louisevanderlith/xchange/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(poxy resins.Epoxi) {
	//Buy
	buyCtrl := &controllers.BuyController{}
	buyGroup := routing.NewRouteGroup("buy", mix.JSON)
	buyGroup.AddRoute("Purchase Something", "", "POST", roletype.User, buyCtrl.Post)
	poxy.AddGroup(buyGroup)

	//Balance
	balCtrl := &controllers.BalanceController{}
	balGroup := routing.NewRouteGroup("balance", mix.JSON)
	balGroup.AddRoute("View Balances", "/", "GET", roletype.User, balCtrl.Get)
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
