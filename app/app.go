package app

import (
	"log"
	"net/http"
)

func StartApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/customers", getAllCustomers)
	log.Fatal(
		http.ListenAndServe(":8080", mux),
	)
}
