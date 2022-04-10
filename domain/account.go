package domain

import (
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponse() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
