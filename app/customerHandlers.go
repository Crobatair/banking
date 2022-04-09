package app

import (
	"encoding/json"
	"github.com/crobatair/banking/domain"
	"github.com/crobatair/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

// This is a receiver function, it will bind:
// - findAllCustomers  -> to an instantiated CustomerHandlers in app.go
// This bind, provides a service.FindAllCustomers() that is defined on the Stub
func (ch *CustomerHandlers) findAllCustomers(w http.ResponseWriter, r *http.Request) {
	filters := r.URL.Query()
	customers, err := ch.service.FindAllCustomers(filters)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.FindCustomerById(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func findCustomer(w http.ResponseWriter, r *http.Request) {
	customer := domain.Customer{
		Name:    "John",
		City:    "New York",
		Zipcode: "10001",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	customer := domain.Customer{
		Name:    "John",
		City:    "New York",
		Zipcode: "10001",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}
