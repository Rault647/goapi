package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {
	"alex": {
		AuthToken:	"123ABC",
		Username:		"alex",
	},
	"jason": {
		AuthToken:	"456DEF",
		Username:		"jason",
	},
	"marie": {
		AuthToken:	"789GHI",
		Username:		"marie",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"alex": {
		Coins:		100,
		Username:	"alex",
	},
	"jason": {
		Coins:		200,
		Username:	"jason",
	},
	"marie": {
		Coins:		300,
		Username:	"marie",
	},
}

// Looks up data from the mockLoginDetails map above.
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

// Looks up data from the mockCoinDetails map above.
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

// Does nothing (we're not using a real database, so there are no errors from "database")
func (d *mockDB) SetupDatabase() error {
	return nil
}