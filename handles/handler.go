package handles

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() http.Handler {
	/*
	buyCtrl := &handles.Buy{}
		balCtrl := &handles.Balance{}
		e.JoinBundle("/", roletype.User, mix.JSON, buyCtrl, balCtrl)
	*/
	r := mux.NewRouter()



	return r
}
