package main

import (
	"Go-Rest/store"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	port := "3000"
	router := store.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
