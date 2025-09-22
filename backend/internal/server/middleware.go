package server

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Chain(handler http.Handler, middleware ...Middleware) http.Handler {
	for _, m := range middleware {
		handler = m(handler)
	}

	return handler
}
