package middleware

import (
	"net/http"
)

// WEB middlewares goes here

// RequestPathLogger Middleware function, which log the called request path
func RequestPathLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Request Path is : " + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
