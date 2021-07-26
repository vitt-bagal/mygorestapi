package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// calls to handlers functions
	myRouter.HandleFunc("/buy-item/{name}", buyItem).Methods("GET")
	myRouter.HandleFunc("/buy-item-qty/{name}/{quantity:[0-9]+}", buyItemQuantity).Methods("GET")
	myRouter.HandleFunc("/buy-item-qty-price/{name}/{quantity:[0-9]+}/{price}", buyItemQuantityPrice).Methods("GET")
	myRouter.HandleFunc("/show-summary", showsummary).Methods("GET")
	myRouter.HandleFunc("/fast-buy-item/{name}", fastBuyItem).Methods("GET")
	myRouter.HandleFunc("/", homePage)
	// start server on 9091 port
	log.Fatal(http.ListenAndServe(":9091", myRouter))
}
