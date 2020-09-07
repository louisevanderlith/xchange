package core

import "github.com/louisevanderlith/husk"

type context struct {
	//Ledger
	Credits husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{
		Credits: husk.NewTable(Currency{}),
	}
}

func Shutdown() {
	ctx.Credits.Save()
}
