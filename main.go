package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"first_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	http.HandleFunc("/customers", GetAllCustomers)
	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)

}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
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
