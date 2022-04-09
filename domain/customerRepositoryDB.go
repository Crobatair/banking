package domain

import (
	"database/sql"
	"github.com/crobatair/banking/errs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	var c Customer
	row := d.client.QueryRow("SELECT * FROM customers WHERE customer_id = ?", id)
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}

		log.Println("Error scanning the filtered row" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}

	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(f url.Values) ([]Customer, *errs.AppError) {
	findAllSql, queryErr := constructSqlFindAllQuery(f)
	if queryErr != nil {
		return nil, queryErr
	}

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("Error while scanning customer table" + err.Error())
		}
		customers = append(customers, c)
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
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDb{client}
}
