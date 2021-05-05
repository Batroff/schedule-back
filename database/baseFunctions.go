package database

import (
	"context"
	"github.com/batroff/schedule-back/models/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type NilUriError struct{}

func (e *NilUriError) Error() string {
	return "Connection URI is nil!"
}

func disconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Printf("Disconnect error. %s", err)
	}
}

func connect(config *config.MongoConfig) (*mongo.Client, context.Context, error) {
	if config.Host == "" {
		return nil, nil, new(NilUriError)
	}

	// Create new client
	client, clientCreateErr := mongo.NewClient(options.Client().ApplyURI(config.Host))
	if clientCreateErr != nil {
		return nil, nil, errors.Wrap(clientCreateErr, "Client creation error.")
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	connectionErr := client.Connect(ctx)
	if connectionErr != nil {
		return nil, nil, errors.Wrap(connectionErr, "Connection error.")
	}

	return client, ctx, nil
}

func SetConfig(Host string) *config.MongoConfig {
	return &config.MongoConfig{
		Host: Host,
	}
}

func SetQuery(documentName, collectionName string) *config.MongoQuery {
	return &config.MongoQuery{
		DocumentName:   documentName,
		CollectionName: collectionName,
	}
}
