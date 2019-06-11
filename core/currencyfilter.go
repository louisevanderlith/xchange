package core

import (
	"github.com/louisevanderlith/husk"
)

type currencyFilter func(obj *Currency) bool

func (f currencyFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Currency))
}

func byEntity(entityKey husk.Key) currencyFilter {
	return func(obj *Currency) bool {
		return obj.EntityKey == entityKey
	}
}
