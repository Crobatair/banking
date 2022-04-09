package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartApp() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", findAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", findCustomer).Methods(http.MethodGet)
	log.Fatal(
		http.ListenAndServe(":8080", router),
	)
}
