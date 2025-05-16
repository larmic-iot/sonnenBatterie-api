package middleware

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

func MetricsMiddleware(requestCount *prometheus.CounterVec, requestLatency *prometheus.HistogramVec) mux.MiddlewareFunc {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Route-Name aus mux extrahieren
			route := mux.CurrentRoute(r)
			routeName := "unknown"
			if route != nil {
				if name := route.GetName(); name != "" {
					routeName = name
				}
			}

			timer := prometheus.NewTimer(requestLatency.WithLabelValues(r.Method, routeName))
			defer timer.ObserveDuration()

			requestCount.WithLabelValues(r.Method, r.URL.Path).Inc()

			inner.ServeHTTP(w, r)
		})
	}
}
