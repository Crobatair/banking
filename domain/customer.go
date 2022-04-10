package domain

import (
	"github.com/crobatair/banking/dto"
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

func (c Customer) statusAsTest() string {
	if c.Status == "1" {
		return "active"
	} else {
		return "inactive"
	}
}
func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsTest(),
	}
}

type CustomerRepository interface {
	FindAll(url.Values) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
