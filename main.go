package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simple-account/api/resources"
	"simple-account/configuration"
	"simple-account/middleware/db_middleware"
	"simple-account/middleware/input_output_middleware"
)

func main() {

	log.Print("Starting application")
	router := mux.NewRouter().StrictSlash(true)

	// Middlewares
	router.Use(input_output_middleware.InputOutputMiddleware)
	router.Use(db_middleware.DBMiddleware)


	// Routes
	router.HandleFunc("/account", resources.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/account", resources.GetAccountList).Methods(http.MethodGet)
	router.HandleFunc("/account/{id}", resources.GetSingleAccount).Methods(http.MethodGet)

	log.Print("Routes loaded")
	log.Fatal(http.ListenAndServe(*configuration.GetServerPort(), router))
}

