package middleware

import (
	"net/http"
	"time"
)

const requestTimeoutDuration = time.Second * 30

func TimeoutMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.TimeoutHandler(h, requestTimeoutDuration, "request timeout")
		h.ServeHTTP(w, r)
	})
}
