package middleware

import "net/http"

func Authentication(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		provided_token := r.Header.Get("Authentication-Token")
		if provided_token == "" {
			http.Error(w, "Authentication-Token header required", http.StatusBadRequest)
		} else if provided_token != token {
			http.Error(w, "Invalid Authentication-Token", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
