package middleware

import (
	"errors"
	"net/http"

	"github.com/Rault647/goapi/api"
	"github.com/Rault647/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid username or token.")

// Takes in and returns an http handler interface.
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {	// HandlerFunc taken from http package
		// Logic for authorizing the http request
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// Check if the username and auth token is correct using database
		var database *tools.DatabaseInterface	// Pointer to database using an interface type
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		// If client wasn't found with the username or the token doesn't match.
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// Calls the next middleware in line or the Handler function
		// for the endpoint if there is no middleware left.
		// Middleware1 -> next.ServeHTTP -> Middleware2 -> ... -> HandlerFunc
		next.ServeHTTP(w, r)
	})
}
