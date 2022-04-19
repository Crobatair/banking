package domain

import (
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/errs"
	"strconv"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponse() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (a Account) AmountAsString() string {
	return strconv.FormatFloat(a.Amount, 'f', 2, 64)
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindByAccountId(string) (*Account, *errs.AppError)
	UpdateBalance(string, float64) *errs.AppError
}
