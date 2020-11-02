package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	connectionTimeout = 1 * time.Second
)

type Client struct {
	*mongo.Client
}

func NewClient(ctx context.Context) (*mongo.Client, error) {
	clientContext, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	client, err := connect(clientContext)
	if err != nil {
		return nil, err
	}

	if err := ping(ctx, client); err != nil {
		return nil, errors.Wrap(err, "ping to connection error")
	}

	return client, nil
}

func connect(ctx context.Context) (*mongo.Client, error) {
	uri := fmt.Sprintf("%s://%s:%s@%s%s", "mongodb", "root", "tvisodb", "localhost:27017", "")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("connection error: %s", uri))
	}

	if err := ping(ctx, client); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("ping to connection error: %s", uri))
	}

	return client, nil
}

func ping(ctx context.Context, client *mongo.Client) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "ping to connection error")
	}

	return nil
}
