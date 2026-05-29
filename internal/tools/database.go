package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections
type LoginDetails struct {
	AuthToken string
	Username	string
}

type CoinDetails struct {
	Coins			int64
	Username	string
}

// Defines methods required for the api
// Using an interface allows us to swap databases easily, as long as
// GetUserLoginDetails, GetUserCoins, and SetupDatabase methods are defined.
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}	// Mock database found in internal/tools/mockdb.go
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}