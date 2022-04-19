package domain

import (
	"github.com/crobatair/banking/errs"
	"github.com/crobatair/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) UpdateBalance(accountId string, amount float64) *errs.AppError {
	query := "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
	_, err := d.client.Exec(query, amount, accountId)
	if err != nil {
		logger.Error("Error updating balance" + err.Error())
		return errs.NewUnexpectedError("Error updating balance" + err.Error() + ", the transaction will be reverted")
	}

	return nil
}

func (d AccountRepositoryDb) FindByAccountId(s string) (*Account, *errs.AppError) {
	var a Account
	sqlFind := `SELECT * FROM accounts WHERE account_id = ?`
	err := d.client.Get(&a, sqlFind, s)
	if err != nil {
		logger.Error("Error while finding account by account_id" + err.Error())
		return nil, errs.NewNotFoundError("Account not found")
	}
	return &a, nil
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?,?,?,?,?)"
	res, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account." + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database.")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id." + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database.")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepository {
	return &AccountRepositoryDb{client}
}
