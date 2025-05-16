package api

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		Name:    "Metrics",
		Method:  "GET",
		Pattern: "/metrics",
		HandlerFunc: func(e env.Environment, w http.ResponseWriter, r *http.Request) {
			promhttp.Handler().ServeHTTP(w, r)
		},
	},
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/api/",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "Consumption",
		Method:      "GET",
		Pattern:     "/api/consumption",
		HandlerFunc: ConsumptionHandler,
	},
	Route{
		Name:        "System",
		Method:      "GET",
		Pattern:     "/api/system",
		HandlerFunc: SystemHandler,
	},
	Route{
		Name:        "Status",
		Method:      "GET",
		Pattern:     "/api/status",
		HandlerFunc: StatusHandler,
	},
}
