package app

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"first_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		Customer{
			Name:    "John",
			City:    "New York",
			Zipcode: "10001",
		},
		Customer{
			Name:    "Jane",
			City:    "New York",
			Zipcode: "10001",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}
