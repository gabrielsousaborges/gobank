package main

import (
	"gobank/domains/transaction/transport"
	"net/http"
)

func main() {

	http.HandleFunc("/transaction", transport.NewTransactionHadler)

	http.HandleFunc("/createUser", transport.NewUserHandler)

	http.ListenAndServe(":8080", nil)

}
