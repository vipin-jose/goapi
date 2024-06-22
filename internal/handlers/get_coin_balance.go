package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/vipin-jose/goapi/api"
	"github.com/vipin-jose/goapi/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params api.CoinBalanceParams
	var err error
	err = schema.NewDecoder().Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var balance int64
	balance = (*database).GetUserCoins(params.Username).Coins
	// if err != nil {
	// 	api.InternalErrorHandler(w)
	// 	return
	// }

	var response api.CoinBalanceResponse
	response.Code = 200
	response.Balance = balance
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
