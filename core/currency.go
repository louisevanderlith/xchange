package core

import (
	"errors"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

const unitValue = 60

//Currency is the base credit used to "purchase" assets
type Currency struct {
	HeroKey  hsk.Key
	Quantity int64
}

func (c Currency) Valid() error {
	return validation.Struct(c)
}

//GetBalance returns the hero's token balance
func GetBalance(heroKey hsk.Key) (int64, error) {
	if heroKey == nil {
		return 0, errors.New("invalid key")
	}

	coll, err := ctx.Credits.Find(1, 999, byHero(heroKey))

	if err != nil {
		return 0, err
	}

	if !coll.Any() {
		return 0, nil
	}

	itor := coll.GetEnumerator()

	var sum int64
	for itor.MoveNext() {
		rec := itor.Current().(hsk.Record)
		curr := rec.GetValue().(Currency)

		sum += curr.Quantity * unitValue
	}

	return sum, nil
}
