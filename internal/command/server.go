package command

import (
	"fmt"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
	"github.com/spf13/cobra"
	"os"
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
		fmt.Fprint(os.Stderr, "empty path\n")
		os.Exit(3)
	}

	e.cfg.Server.Path = args[0]

	openapi, err := openapi3.Parse(e.cfg.Server.Path)
	if err != nil {
		fmt.Fprint(os.Stderr, "can't parse yaml\n")
		os.Exit(3)
	}

	s := server.NewServer(e.cfg.Server, openapi)

	if err := s.Run(); err != nil {
		fmt.Fprint(os.Stderr, "server error\n")
		os.Exit(3)
	}
}
