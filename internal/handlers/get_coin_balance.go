package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rault647/goapi/api"
	"github.com/Rault647/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	/*
		Assuming the call has already ran through the Authorization middleware,
		so we just need to grab the username from the parameters passed in.
		We can do this by decoding the parameters to the CoinBalanceParams struct.
	*/
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	
	// Grab the parameters in the URL and set them to the values in the struct.
	err = decoder.Decode(&params, r.URL.Query())

	if err !=  nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()	// Instantiate a database interface.
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)	// Call GetUserCoins from the instantiated database interface.
	if tokenDetails = nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Set the value to the response struct
	var response = api.CoinBalanceResponse {
		Balance:	(*tokenDetails).Coins,
		Code:			http.StatusOK,
	}

	// Write the response to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
