package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/exitcode"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
)

const version = "0.0.0"

func main() {
	run()
}

func run() {
	flag.Usage = func() {
		fmt.Println(`usage: dummy [flags] [path]

- server [path] - run mock server
- version - show version and exit`)
	}

	flag.Parse()

	cfg := config.NewConfig()

	flag.StringVar(&cfg.Server.Port, "port", "8080", "")
	flag.StringVar(&cfg.Logger.Level, "logger-level", "INFO", "")

	if len(flag.Args()) == 2 && flag.Args()[0] == "server" {
		cfg.Server.Path = flag.Args()[1]

		openapi, err := openapi3.Parse(cfg.Server.Path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "specification parse error: %v\n", err)
			os.Exit(exitcode.Failure)
		}

		l := logger.NewLogger(cfg.Logger.Level)
		h := server.NewHandlers(openapi, l)
		s := server.NewServer(cfg.Server, l, h)

		err = s.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "server run error: %v\n", err)
			os.Exit(exitcode.Failure)
		}
	}

	if len(flag.Args()) > 0 && flag.Args()[0] == "version" {
		fmt.Println(version)
		os.Exit(exitcode.Success)
	}
}
