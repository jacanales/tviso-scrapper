package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/mongo"

	"tviso-scrapper/pkg/platform/mongodb"
	"tviso-scrapper/pkg/tviso"
)

const (
	database   = "collections"
	collection = "tviso"
)
type MongoDB struct {
	client *mongo.Client
	encoder *bson.Encoder
}

func NewMongoDBClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cli, err := mongodb.NewClient(ctx)
	if err != nil {
		log.Panicf("cannot connect to mongodb")
	}

	return cli
}

func NewMongoDBRepository() tviso.WriteRepository {
	cli := NewMongoDBClient()

	got := make(bsonrw.SliceWriter, 0, 1024)
	vw, err := bsonrw.NewBSONValueWriter(&got)
	if err != nil {
		log.Panicf("cannot connect to mongodb")
	}
	encoder, err := bson.NewEncoder(vw)

	return MongoDB{
		client: cli,
		encoder: encoder,
	}
}

func (m MongoDB) StoreCollection(media []tviso.Media) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	for _, md := range media {
		err := m.encoder.Encode(md)
		if err != nil {
			return err
		}

		_, err = m.client.Database(database).Collection(collection).InsertOne(ctx, md)
		if err != nil {
			return err
		}
	}

	return nil
}
