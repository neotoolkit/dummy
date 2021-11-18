package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/go-dummy/dummy/internal/exitcode"
)

func (e *Executor) initRoot() {
	rootCmd := &cobra.Command{
		Use:   "dummy",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				fmt.Fprintln(os.Stdout, "Usage: dummy")
				os.Exit(exitcode.Success)
			}
		},
	}

	e.rootCmd = rootCmd
}
