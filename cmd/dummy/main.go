package main

import (
	"fmt"
	"os"

	"github.com/go-dummy/dummy/internal/command"
	"github.com/go-dummy/dummy/internal/exitcode"
)

const version = "0.0.0"

func main() {
	e := command.NewExecutor(version)

	if err := e.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed executing command with error %v\n", err)
		os.Exit(exitcode.Failure)
	}
}
