package routes

import (
	"log"
	"net/http"
)

// MiddlewareFunc is a custom Middleware type
type MiddlewareFunc func(http.Handler) http.Handler

// RequestLogger logs each request
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI, r.RemoteAddr, r.UserAgent())
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// HeaderBinding checks if HTTP Headers are set correctly and sets response Headers
func HeaderBinding(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			if r.Header.Get("Content-Type") != "application/json" {
				log.Fatal("Wrong Content-Type in POST request. application/json expected")
			}
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
