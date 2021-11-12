package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (e *Executor) initRoot() {
	rootCmd := &cobra.Command{
		Use:   "dummy",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				fmt.Println("Usage: dummy")
			}
		},
	}

	e.rootCmd = rootCmd
}
