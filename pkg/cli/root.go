package cli

import "github.com/spf13/cobra"

// CobraFn function definition of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var rootCmd = &cobra.Command{
	Use: "tviso-cli",
}

func Execute() error {
	rootCmd.AddCommand(InitCollectionListCmd())

	return rootCmd.Execute()
}
