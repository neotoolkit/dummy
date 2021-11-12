package command

import "github.com/spf13/cobra"

type Executor struct {
	rootCmd   *cobra.Command
	serverCmd *cobra.Command
}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Execute() error {
	return e.rootCmd.Execute()
}
