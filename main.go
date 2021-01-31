package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"sonnen-batterie-api/api"
	"sonnen-batterie-api/api/env"
)

func main() {
	log.Println("Hello sonnenBatterie-api!")

	environment, err := env.ReadVariables()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(environment)

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
