package db

import (
	"context"
	"github.com/Sakenzhassulan/user-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Collection *mongo.Collection
}

func NewDBCollection(config *config.Config) (*DB, error) {
	clientOptions := options.Client().ApplyURI(config.DbUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database(config.DbName).Collection(config.DbCollection)
	return &DB{
		Collection: collection,
	}, nil
}
