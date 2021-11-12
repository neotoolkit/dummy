package main

import (
	"fmt"
	"os"

	"github.com/go-dummy/dummy/internal/command"
)

func main() {
	e := command.NewExecutor()

	if err := e.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed executing command with error %v\n", err)
		os.Exit(3)
	}
}
