package middleware

import (
	"github.com/leandroyou/http-car/types"
	"io"
	"net/http"
)

func init() {
	registerMiddleware(authMiddleware)
}

// authMiddleware is an example of a middleware layer. It handles the request authorization
// by checking for a key in the url.
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestKey := types.APIkey(r.URL.Query().Get("key"))
		if len(requestKey) == 0 || !requestKey.IsValid() {
			// Report Unauthorized
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"invalid_key"}`)
			return
		}

		next.ServeHTTP(w, r)
	})
}
