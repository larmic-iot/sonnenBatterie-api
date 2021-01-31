package api

import (
	"log"
	"net/http"
	"sonnen-batterie-api/api/env"
	"time"
)

func Logger(e env.Environment, inner HandlerFunc, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(e, w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
