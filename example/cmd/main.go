package main

import (
	"fmt"
	"os"

	centricerrorwrapper "github.com/pixel8labs/errorwrapper"
	cmd "github.com/pixel8labs/errorwrapper/example"
)

func main() {
	// Init.
	err := centricerrorwrapper.NewCentralizeErrors()
	if err != nil {
		fmt.Println("err when initialize central-error-wrapper", err)
		return
	}

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <target>")
		return
	}

	switch args[0] {
	case "app":
		cmd.InitApp()
	case "api":
		cmd.InitApi()
	}
}
