package middleware

import (
	"fmt"
	"net/http"
)

func init() {
	registerMiddleware(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("mensagem de teste")

		next.ServeHTTP(w, r)
	})
}
