package cli

import (
	"github.com/spf13/cobra"

	`tviso-scrapper/pkg/tviso`
	"tviso-scrapper/pkg/tviso/repository"
)

func InitCollectionListCmd() *cobra.Command {
	collectionListCmd := &cobra.Command{
		Use:   "list",
		Short: "List user collection",
		RunE:  getCollectionListFn(),
	}

	return collectionListCmd
}

func getCollectionListFn() CobraFnE {
	return func(cmd *cobra.Command, args []string) error {
		return tviso.GetUserCollection(
			repository.NewTvisoAPI(),
			repository.NewStdOut(),
		)
	}
}
