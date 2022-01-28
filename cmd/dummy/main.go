package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cristalhq/acmd"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/parse"
	"github.com/go-dummy/dummy/internal/server"
)

const version = "0.2.1"

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

				fs := flag.NewFlagSet("dummy", flag.ContinueOnError)
				fs.StringVar(&cfg.Server.Port, "port", "8080", "")
				fs.StringVar(&cfg.Logger.Level, "logger-level", "INFO", "")
				if err := fs.Parse(args[1:]); err != nil {
					return err
				}

				api, err := parse.Parse(cfg.Server.Path)
				if err != nil {
					return fmt.Errorf("specification parse error: %w", err)
				}

				l := logger.NewLogger(cfg.Logger.Level)
				h := server.NewHandlers(api, l)
				s := server.NewServer(cfg.Server, l, h)

				go func() {
					if err := s.Run(); !errors.Is(err, http.ErrServerClosed) {
						l.Logger.Err(err).Msg("run server")
					}
				}()

				interrupt := make(chan os.Signal, 1)
				signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

				x := <-interrupt
				l.Info().Msgf("received `%v`", x)

				const timeout = 5 * time.Second

				ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
				defer cancelFunc()

				return s.Stop(ctx)
			},
		},
	}

	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName: "Dummy",
		Version: version,
	})

	return r.Run()
}
