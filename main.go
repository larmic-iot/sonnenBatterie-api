package main

import (
	"log"
	"net/http"
	"sonnen-batterie-api/api"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Hello sonnenBatterie-api!")

	tokens, err := api.ReadAuthenticationTokens()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(tokens)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", HelloHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello sonnenBatterie-api!"))
}
