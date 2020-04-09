package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func InitCollectionListCmd() *cobra.Command {
	collectionListCmd := &cobra.Command{
		Use: "list",
		Short: "List user collection",
		Run: getCollectionListFn(),
	}

	collectionListCmd.Flags().StringP("writer", "w", "csv", "Select writer")

	return collectionListCmd
}

func getCollectionListFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("list collection")
	}
}