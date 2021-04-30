package database

import (
	"context"
	"github.com/batroff/schedule-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InsertMany(dbName, collectionName string, groups *[]models.Group) error {
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

func InsertGroupList(dbName, collectionName string) error {
	client, ctx := connect("mongodb://localhost:27017")
	defer disconnect(client, ctx)

	groupList := models.CreateGroupList()

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	_, err := collection.InsertOne(ctx, groupList)

	if err != nil {
		return err
	}

	return nil
}

func GetGroupList(dbName, collectionName string) (models.GroupList, error) {
	var result = models.GroupList{}
	var err error = nil

	client, ctx := connect("mongodb://localhost:27017")
	defer disconnect(client, ctx)

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	err = collection.FindOne(context.Background(), bson.D{}).Decode(&result)

	//cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("%v", err)
		return models.GroupList{}, err
	}

	return result, nil
}

func FindGroup(dbName, collectionName, groupName string, subgroup string) (models.Group, error) {
	var group models.Group
	var err error = nil

	client, ctx := connect("mongodb://localhost:27017")
	defer disconnect(client, ctx)

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

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

func connect(URI string) (*mongo.Client, context.Context) {
	if URI == "" {
		log.Panicf("Connection URI is empty!")
	}

	// Create new client
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
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
