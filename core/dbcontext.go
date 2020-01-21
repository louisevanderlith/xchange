package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	//Ledger
	Credits husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Credits: husk.NewTable(Currency{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Credits.Save()
}
