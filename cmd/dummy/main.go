package main

import (
	"fmt"
	"os"

	"github.com/go-dummy/dummy/internal/command"
	"github.com/go-dummy/dummy/internal/exitcode"
)

func main() {
	e := command.NewExecutor()

	if err := e.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed executing command with error %v\n", err)
		os.Exit(exitcode.Failure)
	}
}
