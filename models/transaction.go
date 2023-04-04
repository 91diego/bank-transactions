package models

type TransactionResponse struct {
	Message        string   `json:"message"`
	SummaryBalance *Balance `json:"summary_balance"`
}

type Transaction struct {
	ID                string `json:"id" csv:"id"`
	TransactionDate   string `json:"transactionDate" csv:"transaction_date"`
	TransactionAmount string `json:"transactionAmount" csv:"transaction_amount"`
	TransactionType   string `json:"transactionType" csv:"transaction_type"`
	BlanceID          string `json:"balanceID"`
}
