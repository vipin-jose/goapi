package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]*LoginDetails{
	"vipin": {
		Username:  "vipin",
		AuthToken: "1234",
	},
	"sayana": {
		Username:  "sayana",
		AuthToken: "5678",
	},
}

var mockCoinDetails = map[string]*CoinDetails{
	"vipin": {
		Coins:    100,
		Username: "vipin",
	},
	"sayana": {
		Coins:    200,
		Username: "sayana",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(2 * time.Second)
	return mockLoginDetails[username]
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(2 * time.Second)
	return mockCoinDetails[username]
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
