package domain

import (
	"database/sql"
	"fmt"
	"github.com/crobatair/banking/errs"
	"github.com/crobatair/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/url"
	"os"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	var c Customer
	err := d.client.Get(&c, "SELECT * FROM customers WHERE customer_id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error scanning the filtered row" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}
	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(f url.Values) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	findAllSql, queryErr := constructSqlFindAllQuery(f)
	if queryErr != nil {
		return nil, queryErr
	}

	err := d.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("Error scanning the filtered row" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}
	return customers, nil
}

func constructSqlFindAllQuery(f url.Values) (string, *errs.AppError) {
	query := "SELECT * from customers"
	if f.Get("status") != "" {
		status := f.Get("status")
		statusParam := "0"
		if status == "active" {
			statusParam = "1"
		} else if status == "inactive" {
			statusParam = "0"
		} else {
			return query, errs.NewBadRequestError("Invalid value for filter status")
		}
		query += " WHERE status = " + statusParam

	}
	return query, nil
}

func NewCustomerRepositoryDb() CustomerRepository {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDb{client}
}
