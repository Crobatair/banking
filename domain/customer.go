package domain

import (
	"github.com/crobatair/banking/errs"
	"net/url"
)

type Customer struct {
	Id          string `db:"customer_id" json:"id"`
	Name        string `json:"first_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zip_code"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(url.Values) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
