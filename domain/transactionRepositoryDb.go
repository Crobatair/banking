package domain

import (
	"github.com/crobatair/banking/dto"
	"github.com/crobatair/banking/errs"
	"github.com/crobatair/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDb) SaveTransaction(r *dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	var transaction dto.TransactionResponse
	sqlTransaction := "INSERT INTO transactions (account_id, amount, transaction_type) VALUES (?, ?, ?)"
	res, err := d.client.Exec(sqlTransaction, r.Account, r.Amount, r.TransactionType)
	if err != nil {
		logger.Error("Error while saving transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Error while saving transaction" + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id." + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database.")
	}

	transaction.TransactionId = strconv.FormatInt(id, 10)

	return &transaction, nil
}

func (d TransactionRepositoryDb) RevertTransaction(transactionId string) *errs.AppError {
	sqlDelete := "DELETE FROM transactions WHERE transaction_id = ?"
	_, err := d.client.Exec(sqlDelete, transactionId)
	if err != nil {
		logger.Error("Error while reverting transaction" + err.Error())
		return errs.NewUnexpectedError("Error while reverting transaction" + err.Error())
	}
	return nil
}

func NewTransactionRepositoryDb(client *sqlx.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{client: client}
}
