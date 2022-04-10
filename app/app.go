package app

import (
	"fmt"
	"github.com/crobatair/banking/domain"
	"github.com/crobatair/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("A defined environment variable is missing")
	}
}

func StartApp() {
	sanityCheck()
	router := mux.NewRouter()

	// Instance a new handler, that will take the **CustomerRepositoryStub** and build a new **CustomerService**
	// This customer service, implements the **CustomerService** interface
	// The **CustomerService** interface is used to interact with the **CustomerRepositoryStub**
	// The **CustomerRepositoryStub**, it's an abstraction of a remote service / database, but assumes that a default
	// implementation is available and can provide the data
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/api/customers", ch.findAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	s := os.Getenv("SERVER_ADDRESS")
	p := os.Getenv("SERVER_PORT")

	log.Fatal(
		http.ListenAndServe(fmt.Sprintf("%s:%s", s, p), router),
	)
}
