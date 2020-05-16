package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"

	"tviso-scrapper/pkg/platform/mongodb"
)

func InitCreateMongoDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create_database {dbname}",
		Short: "Create a new MongoDB database from scratch",
		Args: cobra.MinimumNArgs(1),
		RunE: createDatabaseFn(),
	}

	return cmd
}

func createDatabaseFn() CobraFnE {
	return func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		dbname := args[0]

		cli, err := mongodb.NewClient(ctx)
		if err != nil {
			return err
		}
		defer func () {
			_ = cli.Disconnect(ctx)
		}()

		collection := cli.Database(dbname).Collection("delete_me")
		res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		if err != nil {
			return err
		}

		dRes, err := collection.DeleteOne(ctx, bson.D{{"_id", res.InsertedID}})
		if err != nil {
			return err
		}
		fmt.Println(res.InsertedID)
		fmt.Println(dRes)

		return nil
	}
}