package models

type Balance struct {
	ID            string  `json:"id"`
	Total         float64 `json:"total"`
	DebitAvarage  float64 `json:"debitAvarage"`
	CreditAvarage float64 `json:"creditAvarage"`
	Transactions  string  `json:"transactions"`
	UserID        string  `json:"userID"`
}
