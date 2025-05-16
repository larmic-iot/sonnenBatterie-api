package main

import (
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"sonnen-batterie-api/api"
	"sonnen-batterie-api/api/env"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func main() {
	log.Println("Hello sonnenBatterie-api!")

	environment, err := env.ReadVariables()

	if err != nil {
		log.Fatal(err)
	}

	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestLatency)

	router := api.NewRouter(environment, requestCount, requestLatency)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
