package command

import (
	"github.com/spf13/cobra"
)

func (e *Executor) initServer() {
	e.serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start dummy server",
		Run:   e.executeServer,
	}

	e.rootCmd.AddCommand(e.serverCmd)
}

func (e *Executor) executeServer(_ *cobra.Command, args []string) {}
