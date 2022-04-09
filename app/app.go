package app

import (
	"github.com/crobatair/banking/domain"
	"github.com/crobatair/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartApp() {
	router := mux.NewRouter()

	// Instance a new handler, that will take the **CustomerRepositoryStub** and build a new **CustomerService**
	// This customer service, implements the **CustomerService** interface
	// The **CustomerService** interface is used to interact with the **CustomerRepositoryStub**
	// The **CustomerRepositoryStub**, it's an abstraction of a remote service / database, but assumes that a default
	// implementation is available and can provide the data
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/api/customers", ch.findAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", findCustomer).Methods(http.MethodGet)
	log.Fatal(
		http.ListenAndServe(":8080", router),
	)
}
