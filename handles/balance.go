package handles

import (
	"net/http"
)

//BalanceController can only display the logged in user's balance
func GetBalance(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
}
