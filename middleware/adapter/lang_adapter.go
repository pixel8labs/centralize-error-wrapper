package adapter

import (
	"net/http"

	centricerrorwrapper "github.com/pixel8labs/errorwrapper"
)

func LangAdapterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := defaultLang

		if r.Header.Get(httpHeaderAcceptLang) != defaultString {
			lang = r.Header.Get(httpHeaderAcceptLang)
		}
		if r.Header.Get(httpHeaderAcceptLanguage) != defaultString {
			lang = r.Header.Get(httpHeaderAcceptLanguage)
		}

		// Set lang.
		centricerrorwrapper.SetLang(lang)
		next.ServeHTTP(w, r)
	})
}
