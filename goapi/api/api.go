package api

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
	// Error code
	Code int

	// Error message
	Message string
}