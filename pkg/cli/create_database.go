package cli

import (
	"context"
	"fmt"
	"time"
	"tviso-scrapper/pkg/platform/mongodb"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	timeout = 5 * time.Second
	pi      = 3.14159
)

func InitCreateMongoDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create_database {dbname}",
		Short: "Create a new MongoDB database from scratch",
		Args:  cobra.MinimumNArgs(1),
		RunE:  createDatabaseFn(),
	}

	return cmd
}

func createDatabaseFn() CobraFnE {
	return func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		dbname := args[0]

		cli, err := mongodb.NewClient(ctx)
		if err != nil {
			return err
		}

		defer func() {
			_ = cli.Disconnect(ctx)
		}()

		collection := cli.Database(dbname).Collection("delete_me")

		res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": pi})
		if err != nil {
			return err
		}

		dRes, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: res.InsertedID}})
		if err != nil {
			return err
		}

		fmt.Println(res.InsertedID)
		fmt.Println(dRes)

		return nil
	}
}
