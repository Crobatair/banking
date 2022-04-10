package service

import (
	"github.com/crobatair/banking/domain"
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (d DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := d.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponse()
	return &response, nil
}

func NewAccountRepository(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
