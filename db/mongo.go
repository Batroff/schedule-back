package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"schedule/structure"
	"time"
)

func InsertMany(dbName, collectionName string, groups *[]structure.Group) error {
	client, ctx := connect("mongodb://localhost:27017")
	defer disconnect(client, ctx)

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	for _, group := range *groups {
		log.Printf("Group %s upload starting...", group.Name)

		_, err := collection.InsertOne(ctx, group)

		if err != nil {
			return err
		}

		log.Printf("Group %s uploaded", group.Name)
	}

	return nil
}

func disconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Panicf("%v", err)
	}
}

func connect(URI string) (*mongo.Client, context.Context) {
	if URI == "" {
		log.Panicf("Connection URI is empty!")
	}

	// Create new client
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Panicf("%v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	err = client.Connect(ctx)
	if err != nil {
		log.Panicf("%v", err)
	}

	return client, ctx
}
