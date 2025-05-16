package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"sonnen-batterie-api/api/client/middleware"
	"sonnen-batterie-api/api/env"

	"github.com/gorilla/mux"
)

func NewRouter(environment env.Environment, requestCount *prometheus.CounterVec, requestLatency *prometheus.HistogramVec) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = Handle404()

	router.Use(middleware.MetricsMiddleware(requestCount, requestLatency))

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
