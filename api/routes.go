package api

import (
	"net/http"
	"sonnen-batterie-api/api/env"
)

type HandlerFunc func(env.Environment, http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(e env.Environment, w http.ResponseWriter, r *http.Request) {
	f(e, w, r)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/api",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "Consumption",
		Method:      "GET",
		Pattern:     "/api/consumption",
		HandlerFunc: ConsumptionHandler,
	},
}
