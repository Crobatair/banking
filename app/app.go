package app

import (
	"log"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
