package cli

import (
	"github.com/spf13/cobra"
)

// CobraFn function definition of run cobra command.
type CobraFn func(cmd *cobra.Command, args []string)
type CobraFnE func(cmd *cobra.Command, args []string) error

func Execute() error {
	var rootCmd = &cobra.Command{
		Use: "tviso-cli",
		SilenceUsage: true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(InitCollectionListCmd())
	rootCmd.AddCommand(InitCreateMongoDatabaseCmd())

	return rootCmd.Execute()
}
