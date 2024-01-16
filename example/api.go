package example

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	centricerrorwrapper "github.com/pixel8labs/errorwrapper"
	"github.com/pixel8labs/errorwrapper/middleware/adapter"
)

type userHandler struct{}

func InitApi() {
	mux := http.NewServeMux()
	mux.Handle("/users/", adapter.LangAdapterMiddleware(&userHandler{}))

	fmt.Println("Running on port: 8080")
	http.ListenAndServe(":8080", mux)
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Invoke.
	errUnmarshal := centricerrorwrapper.Wrap(errors.New("error happened"), centricerrorwrapper.ErrIDUnmarshall, centricerrorwrapper.ErrOptions{
		"variable": "data2",
	})
	toResp := errUnmarshal.Response()

	// Would be done by every err-wrapper inside each-project.
	errByte, _ := json.Marshal(toResp)
	// Would be done by every err-wrapper inside each-project.
	w.Header().Set("Content-Type", "application/json")

	// Set based on err data.
	w.WriteHeader(errUnmarshal.StatusCode)
	w.Write([]byte(errByte))
}
