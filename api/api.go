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
	//Success Code
	Code int
	
	//Account Balance
	Balance int64
}

// Error Response
type Error struct {
	//Error code
	Code int

	//Error message
	Message string
}
