package example

import (
	"errors"
	"fmt"

	centricerrorwrapper "github.com/pixel8labs/errorwrapper"
)

func InitApp() {
	// Set language.
	centricerrorwrapper.SetLang("id")

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
}
