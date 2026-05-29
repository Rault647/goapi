package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error Code
	Code int

	// Error Message
	Message string
}

// Not called directly. Used within RequestError Handler and InternalErrorHandler variables below.
func writeError(w http.ResponseWriter, message string, code int) {
	res := Error {
		Code:			code,
		Message:	message,
	}

	w.Header().Set("Content-Type", "application/json")	// Set content-type (returning a json)
	w.WriteHeader(code)	// Error code

	json.NewEncoder(w).Encode(res)	// Write the error struct out (what the user gets back as a response)
}

var (
	// Returns a specific error message as a response (for errors caused by the user)
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// Returns a generic error message (for internal errors like bugs)
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)