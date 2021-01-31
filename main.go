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

	tokens, err := env.ReadAuthenticationTokens()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(tokens)

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
