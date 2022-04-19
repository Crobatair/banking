package domain

import (
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/errs"
)

type Transaction struct {
	TransactionId   int     `db:"transaction_id"`
	AccountId       int     `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

type TransactionRepository interface {
	SaveTransaction(*dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	RevertTransaction(string) *errs.AppError
}
