package handles

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Buy struct {
}

func (x *Buy) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

func (req *Buy) Create(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
