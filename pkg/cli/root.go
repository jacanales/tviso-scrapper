package cli

import (
	"github.com/spf13/cobra"
)

// CobraFn function definition of run cobra command.
type CobraFnE func(cmd *cobra.Command, args []string) error

func Execute() error {
	var rootCmd = &cobra.Command{
		Use: "tviso-cli",
	}

	rootCmd.AddCommand(InitCollectionListCmd())

	return rootCmd.Execute()
}
