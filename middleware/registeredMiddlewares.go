package middleware

import "net/http"

var registeredMiddlewares []func(next http.Handler) http.Handler

func GetRegisteredMiddlewares() []func(next http.Handler) http.Handler {
	return registeredMiddlewares
}

func registerMiddleware(middleware func(next http.Handler) http.Handler) {
	registeredMiddlewares = append(registeredMiddlewares, middleware)
}

// ProcessMiddleware (this function) makes adding more than one layer of middleware easy
// by specifying them as a list. It will run the last specified handler first.
func ProcessMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
