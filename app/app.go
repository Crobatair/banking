package app

import (
	"fmt"
	"github.com/crobatair/banking/domain"
	"github.com/crobatair/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
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
	dbClient := getDbClient()

	// Instance a new handler, that will take the **CustomerRepositoryStub** and build a new **CustomerService**
	// This customer service, implements the **CustomerService** interface
	// The **CustomerService** interface is used to interact with the **CustomerRepositoryStub**
	// The **CustomerRepositoryStub**, it's an abstraction of a remote service / database, but assumes that a default
	// implementation is available and can provide the data
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerService := service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))
	accountService := service.NewAccountRepository(domain.NewAccountRepositoryDb(dbClient))
	ts := domain.NewTransactionRepositoryDb(dbClient)
	transactionService := service.NewTransactionService(
		ts,
		&accountService,
	)
	ch := CustomerHandlers{customerService}
	ah := AccountHandler{accountService}
	th := TransactionHandler{transactionService}

	router.HandleFunc("/api/customers", ch.findAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", th.MakeTransaction).Methods(http.MethodPost)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	s := os.Getenv("SERVER_ADDRESS")
	p := os.Getenv("SERVER_PORT")

	log.Fatal(
		http.ListenAndServe(fmt.Sprintf("%s:%s", s, p), router),
	)
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
