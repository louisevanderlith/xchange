package core

import "github.com/louisevanderlith/husk"

type context struct {
	//Ledger
	Credits husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Credits: husk.NewTable(new(Currency)),
	}
}

func Shutdown() {
	ctx.Credits.Save()
}
