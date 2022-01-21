package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/cristalhq/acmd"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
)

const version = "0.1.0"

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "server run error: %v\n", err)
	}
}

func run() error {
	cmds := []acmd.Command{
		{
			Name:        "server",
			Description: "run mock server",
			Do: func(ctx context.Context, args []string) error {
				cfg := config.NewConfig()

				cfg.Server.Path = args[0]

				flag.StringVar(&cfg.Server.Port, "port", "8080", "")
				flag.StringVar(&cfg.Logger.Level, "logger-level", "INFO", "")

				openapi, err := openapi3.Parse(cfg.Server.Path)
				if err != nil {
					return fmt.Errorf("specification parse error: %w\n", err)
				}

				l := logger.NewLogger(cfg.Logger.Level)
				h := server.NewHandlers(openapi, l)
				s := server.NewServer(cfg.Server, l, h)

				return s.Run()
			},
		},
	}

	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName: "Dummy",
		Version: version,
	})

	return r.Run()
}
