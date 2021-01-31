package api

import (
	"sonnen-batterie-api/api/env"

	"github.com/gorilla/mux"
)

func NewRouter(environment env.Environment) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = Handle404()

	for _, route := range routes {
		handler := Logger(environment, route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
