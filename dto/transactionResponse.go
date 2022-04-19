package dto

type TransactionResponse struct {
	TransactionId   string `json:"transaction_id"`
	Balance         string `json:"balance"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
}
