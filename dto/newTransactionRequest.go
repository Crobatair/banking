package dto

import (
	"github.com/crobatair/banking/errs"
)

var TRANSACTION_WITHDRAW string = "withdraw"
var TRANSACTION_DEPOSIT string = "deposit"

type TransactionRequestBody struct {
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

type TransactionRequest struct {
	Account         string
	Balance         float64
	TransactionType string
	Amount          float64
}

func (t TransactionRequest) validate() *errs.AppError {

	if t.Account == "" {
		return errs.NewValidationError("Account is required.")
	}

	if t.TransactionType == "" {
		return errs.NewBadRequestError("Transaction type is required.")
	}

	if t.TransactionType != TRANSACTION_WITHDRAW && t.TransactionType != TRANSACTION_DEPOSIT {
		return errs.NewBadRequestError("Transaction type must be either '" + TRANSACTION_WITHDRAW + "' or '" + TRANSACTION_DEPOSIT + "'.")
	}

	if t.Amount <= 0 {
		return errs.NewBadRequestError("Amount must be greater than 0 to perform transaction.")
	}

	return nil
}

func NewTransactionRequest(a string, balance float64, transactionType string, amount float64) (*TransactionRequest, *errs.AppError) {
	request := TransactionRequest{
		Account:         a,
		Balance:         balance,
		TransactionType: transactionType,
		Amount:          amount,
	}
	err := request.validate()
	if err != nil {
		return nil, err
	}

	return &request, nil
}
