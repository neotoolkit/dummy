package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/go-dummy/dummy/internal/exitcode"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
)

func (e *Executor) initServer() {
	e.serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start dummy server",
		Run:   e.executeServer,
	}

	e.serverCmd.Flags().StringVarP(&e.cfg.Server.Port, "port", "p", "8080", "")

	e.rootCmd.AddCommand(e.serverCmd)
}

func (e *Executor) executeServer(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "empty path")
		os.Exit(exitcode.Failure)
	}

	e.cfg.Server.Path = args[0]

	openapi, err := openapi3.Parse(e.cfg.Server.Path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't parse yaml")
		os.Exit(exitcode.Failure)
	}

	s := server.NewServer(e.cfg.Server, openapi)

	if err := s.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "server error")
		os.Exit(exitcode.Failure)
	}
}
