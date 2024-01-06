package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pixel8labs/errorwrapper/middleware/adapter"
)

func MuxLangMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return adapter.LangAdapterMiddleware(next)
	}
}
