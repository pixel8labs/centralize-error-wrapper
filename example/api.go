package example

import (
	"errors"
	"fmt"
	"net/http"

	centricerrorwrapper "github.com/pixel8labs/errorwrapper"
	"github.com/pixel8labs/errorwrapper/middleware"
)

type userHandler struct{}

func InitApi() {
	mux := http.NewServeMux()
	mux.Handle("/users/", middleware.LangMiddleware(&userHandler{}))

	fmt.Println("Running on port: 8080")
	http.ListenAndServe(":8080", mux)
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Invoke.
	errMarshal := centricerrorwrapper.New(centricerrorwrapper.ErrIDMarshall, centricerrorwrapper.ErrOptions{
		"variable": "data",
	})
	errUnmarshal := centricerrorwrapper.Wrap(errors.New("error happened"), centricerrorwrapper.ErrIDUnmarshall, centricerrorwrapper.ErrOptions{
		"variable": "data2",
	})

	// (Optional) casting function.
	cast := centricerrorwrapper.Cast(errMarshal)

	// Print result.
	fmt.Println(
		centricerrorwrapper.Unwrap(errUnmarshal),
		"unwrap",
	)
	fmt.Printf("%+v\n", cast.GetMessage())
	fmt.Printf("%+v\n", errUnmarshal.GetMessage())

	w.Write([]byte(`{"message": "ok"}`))
}
