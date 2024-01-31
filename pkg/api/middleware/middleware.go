package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the Authorization header is present
		token := r.Header.Get("Authorization")
		if token != "5425861" {
			http.Error(w, "Unauthorized: Missing bearer token or wrong token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
