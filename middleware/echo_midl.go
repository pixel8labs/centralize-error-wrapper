package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pixel8labs/errorwrapper/middleware/adapter"
)

func EchoLangMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Convert echo.Context to http.Handler.
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next(c)
		})

		// Use the base middleware.
		adaptedHandler := adapter.LangAdapterMiddleware(handler)

		// Call the adapted handler.
		adaptedHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
