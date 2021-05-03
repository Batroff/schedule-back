package database

import (
	"context"
	"github.com/batroff/schedule-back/models"
	"github.com/batroff/schedule-back/models/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InsertMany(config *config.MongoConfig, query *config.MongoQuery, groups *[]models.Group) error {
	client, ctx := connect(config)
	defer disconnect(client, ctx)

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

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

func InsertGroupList(config *config.MongoConfig, query *config.MongoQuery) error {
	client, ctx := connect(config)
	defer disconnect(client, ctx)

	groupList := models.CreateGroupList()

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	_, err := collection.InsertOne(ctx, groupList)

	if err != nil {
		return err
	}

	return nil
}

func GetGroupList(config *config.MongoConfig, query *config.MongoQuery) (models.GroupList, error) {
	var result = models.GroupList{}
	var err error = nil

	client, ctx := connect(config)
	defer disconnect(client, ctx)

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	err = collection.FindOne(context.Background(), bson.D{}).Decode(&result)

	//cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("%v", err)
		return models.GroupList{}, err
	}

	return result, nil
}

func FindGroup(config *config.MongoConfig, query *config.MongoQuery, groupName string, subgroup string) (models.Group, error) {
	var group models.Group
	var err error = nil

	client, ctx := connect(config)
	defer disconnect(client, ctx)

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	if subgroup != "" {
		err = collection.FindOne(context.Background(), bson.M{"name": groupName, "subgroup": subgroup[0]}).Decode(&group)
	} else {
		err = collection.FindOne(context.Background(), bson.M{"name": groupName}).Decode(&group)
	}
	if err != nil {
		log.Printf("%v", err)
		return models.Group{}, err
	}

	return group, err
}

func disconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Panicf("%v", err)
	}
}

func connect(config *config.MongoConfig) (*mongo.Client, context.Context) {
	if config.Host == "" {
		log.Panicf("Connection URI is empty!")
	}

	// Create new client
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Host))
	if err != nil {
		log.Panicf("%v", err)
	}

	// TODO: rework context to value
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	err = client.Connect(ctx)
	if err != nil {
		log.Panicf("%v", err)
	}

	return client, ctx
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
