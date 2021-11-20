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
		Short: "API mocking with OpenAPI v3.x",
		Long:  `Dummy is an API mocking with OpenAPI v3.x`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				fmt.Println("Usage: dummy")
				os.Exit(exitcode.Success)
			}
		},
	}

	e.rootCmd = rootCmd
}
