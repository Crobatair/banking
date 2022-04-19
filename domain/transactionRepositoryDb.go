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
	tx, err := d.client.Beginx()
	if err != nil {
		logger.Error("Error while creating transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Error while creating transaction")
	}

	var transaction dto.TransactionResponse
	sqlTransaction := "INSERT INTO transactions (account_id, amount, transaction_type) VALUES (?, ?, ?)"
	res, err := tx.Exec(sqlTransaction, r.Account, r.Amount, r.TransactionType)
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

	queryUpdate := "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
	_, err = tx.Exec(queryUpdate, r.Amount, r.Account)
	if err != nil {
		logger.Error("Error while updating account" + err.Error())
		tx.Rollback()

		return nil, errs.NewUnexpectedError("Error while updating account" + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Error while committing transaction")
	}

	return &transaction, nil
}

func NewTransactionRepositoryDb(client *sqlx.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{client: client}
}
