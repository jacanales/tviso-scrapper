package cli

import (
	"tviso-scrapper/pkg/tviso"
	"tviso-scrapper/pkg/tviso/repository"

	"github.com/spf13/cobra"
)

func InitCollectionListCmd() *cobra.Command {
	collectionListCmd := &cobra.Command{
		Use:   "sync",
		Short: "Synchronize user collection",
		RunE:  getCollectionListFn(),
	}

	return collectionListCmd
}

func getCollectionListFn() CobraFnE {
	return func(cmd *cobra.Command, args []string) error {
		return tviso.GetUserCollection(
			repository.NewTvisoAPI(repository.NewHTTPClient(), repository.NewConfig()),
			repository.NewMongoDBRepository(),
		)
	}
}
