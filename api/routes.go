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
		Pattern:     "/sonnen-api/",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/sonnen-api/api",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "Consumption",
		Method:      "GET",
		Pattern:     "/sonnen-api/api/consumption",
		HandlerFunc: ConsumptionHandler,
	},
	Route{
		Name:        "System",
		Method:      "GET",
		Pattern:     "/sonnen-api/api/system",
		HandlerFunc: SystemHandler,
	},
	Route{
		Name:        "Status",
		Method:      "GET",
		Pattern:     "/sonnen-api/api/status",
		HandlerFunc: StatusHandler,
	},
}
