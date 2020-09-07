package core

import (
	"errors"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

const unitValue = 75

//Currency is the base credit used to "purchase" assets
type Currency struct {
	EntityKey hsk.Key
	Quantity  int64
}

func (c Currency) Valid() error {
	return validation.Struct(c)
}

//GetBalance returns the entity's token balance
func GetBalance(entityKey hsk.Key) (int64, error) {
	if entityKey == nil {
		return 0, errors.New("invalid key")
	}

	coll, err := ctx.Credits.Find(1, 999, byEntity(entityKey))

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
		curr := rec.Data().(Currency)

		sum += curr.Quantity * unitValue
	}

	return sum, nil
}
