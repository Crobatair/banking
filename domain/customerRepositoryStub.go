package domain

// CustomerRepositoryStub is a stub for the customer repository
// Provides a struct, to hold all customers
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll
// It's an Impl of CustomerRepository interface because implements a FindAll method and returns a []Customer, error.
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub
// Just a constructor for the CustomerRepositoryStub struct
// Nothing special about it.
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1",
			Name:        "John",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1980-01-01",
			Status:      "active",
		},
		{
			Id:          "2",
			Name:        "Rob",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1980-01-01",
			Status:      "active",
		},
		{
			Id:          "3",
			Name:        "Sally",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1980-01-01",
			Status:      "active",
		},
	}
	return CustomerRepositoryStub{customers}
}
