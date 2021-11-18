package command

import (
	"github.com/spf13/cobra"

	"github.com/go-dummy/dummy/internal/config"
)

type Executor struct {
	cfg     *config.Config
	version string

	rootCmd    *cobra.Command
	serverCmd  *cobra.Command
	versionCmd *cobra.Command
}

func NewExecutor(version string) *Executor {
	e := &Executor{
		cfg: config.NewConfig(),

		version: version,
	}

	e.initRoot()
	e.initServer()
	e.initVersion()

	return e
}

func (e *Executor) Execute() error {
	return e.rootCmd.Execute()
}
