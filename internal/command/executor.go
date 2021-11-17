package command

import (
	"github.com/go-dummy/dummy/internal/config"
	"github.com/spf13/cobra"
)

type Executor struct {
	cfg       *config.Config
	rootCmd   *cobra.Command
	serverCmd *cobra.Command
}

func NewExecutor() *Executor {
	e := &Executor{
		cfg: config.NewConfig(),
	}

	e.initRoot()
	e.initServer()

	return e
}

func (e *Executor) Execute() error {
	return e.rootCmd.Execute()
}
