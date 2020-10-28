package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type currencyFilter func(obj Currency) bool

func (f currencyFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Currency))
}

func byHero(heroKey hsk.Key) currencyFilter {
	return func(obj Currency) bool {
		return obj.HeroKey == heroKey
	}
}
