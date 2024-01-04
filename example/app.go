package main

import (
	"context"
	"errors"
	"fmt"

	centricerrorwrapper "github.com/pixel8labs/errorwrapper"
)

func main() {
	// Init.
	err := centricerrorwrapper.NewCentralizeErrors()
	if err != nil {
		fmt.Println("err when initialize central-error-wrapper", err)
		return
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "Accept-Lang", "en")

	// Invoke.
	errMarshal := centricerrorwrapper.New(ctx, centricerrorwrapper.ErrIDMarshall, centricerrorwrapper.ErrOptions{
		"variable": "data",
	})
	errUnmarshal := centricerrorwrapper.Wrap(ctx, errors.New("error happened"), centricerrorwrapper.ErrIDMarshall, centricerrorwrapper.ErrOptions{
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
	fmt.Printf("%+v\n", errUnmarshal.GetErr())
}
