package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (e *Executor) initVersion() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Version",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println(e.version)
		},
	}

	e.rootCmd.AddCommand(versionCmd)
}
