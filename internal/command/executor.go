package command

import "github.com/spf13/cobra"

type Executor struct {
	rootCmd   *cobra.Command
	serverCmd *cobra.Command
}

func NewExecutor() *Executor {
	e := &Executor{}

	e.initRoot()
	e.initServer()

	return e
}

func (e *Executor) Execute() error {
	return e.rootCmd.Execute()
}
