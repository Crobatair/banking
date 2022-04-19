package service

import (
	"github.com/crobatair/banking/domain"
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/errs"
)

type TransactionService interface {
	MakeTransaction(*dto.TransactionRequestBody, string) (*dto.TransactionResponse, *errs.AppError)
	Withdraw(*dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	Deposit(*dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo           *domain.TransactionRepositoryDb
	accountService *DefaultAccountService
}

func NewTransactionService(repo *domain.TransactionRepositoryDb, accountService *DefaultAccountService) *DefaultTransactionService {
	return &DefaultTransactionService{
		repo:           repo,
		accountService: accountService,
	}
}

// Implementations

func (d DefaultTransactionService) MakeTransaction(r *dto.TransactionRequestBody, accountId string) (*dto.TransactionResponse, *errs.AppError) {
	account, err := d.accountService.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	request, err := dto.NewTransactionRequest(account.AccountId, account.Amount, r.TransactionType, r.Amount)
	if err != nil {
		return nil, err
	}

	switch request.TransactionType {
	case dto.TRANSACTION_DEPOSIT:
		return d.Deposit(request)

	case dto.TRANSACTION_WITHDRAW:
		return d.Withdraw(request)

	default:
		return nil, errs.NewBadRequestError("Invalid transaction type")
	}

}

func (d DefaultTransactionService) Deposit(r *dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	result, err := d.repo.SaveTransaction(r)
	if err != nil {
		return nil, err
	}

	account, _ := d.accountService.FindAccount(r.Account)
	result.Balance = account.AmountAsString()
	return result, nil
}

func (d DefaultTransactionService) Withdraw(r *dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	if r.Amount > r.Balance {
		return nil, errs.NewBadRequestError("Insufficient funds")
	}

	r.Amount *= -1 // make it negative for withdraw

	result, err := d.repo.SaveTransaction(r)
	if err != nil {
		return nil, err
	}

	account, _ := d.accountService.FindAccount(r.Account)
	result.Balance = account.AmountAsString()
	return result, nil
}
