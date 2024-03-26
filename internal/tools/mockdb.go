package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"bob": {
		AuthToken: "456DEF",
		Username:  "bob",
	},
	"alice": {
		AuthToken: "789GHI",
		Username:  "alice",
	},
	// Add more entries here
	"charlie": {
		AuthToken: "XYZ456",
		Username:  "charlie",
	},
	"diana": {
		AuthToken: "ABC789",
		Username:  "diana",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"bob": {
		Coins:    150,
		Username: "bob",
	},
	"alice": {
		Coins:    200,
		Username: "alice",
	},
	"charlie": {
		Coins:    50,
		Username: "charlie",
	},
	"diana": {
		Coins:    75,
		Username: "diana",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
