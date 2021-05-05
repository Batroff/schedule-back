package database

import (
	"context"
	"github.com/batroff/schedule-back/models"
	"github.com/batroff/schedule-back/models/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertGroupList(config *config.MongoConfig, query *config.MongoQuery) error {
	client, ctx, connectErr := connect(config)
	if connectErr != nil {
		return connectErr
	}
	defer disconnect(client, ctx)

	groupList := models.CreateGroupList()

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	_, err := collection.InsertOne(ctx, groupList)
	if err != nil {
		return errors.Wrap(err, "Can not insert group list")
	}

	return nil
}

func InsertGroups(config *config.MongoConfig, query *config.MongoQuery, groups *[]models.Group) error {
	client, ctx, connectErr := connect(config)
	if connectErr != nil {
		return connectErr
	}
	defer disconnect(client, ctx)

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	for _, group := range *groups {
		_, err := collection.InsertOne(ctx, group)
		if err != nil {
			return errors.Wrap(err, "Can not insert groups in db")
		}

		log.Printf("Group %s uploaded", group.Name)
	}

	return nil
}

func FindGroup(config *config.MongoConfig, query *config.MongoQuery, groupName string, subgroup string) (*models.Group, error) {
	var group *models.Group

	client, ctx, connectErr := connect(config)
	if connectErr != nil {
		return nil, connectErr
	}
	defer disconnect(client, ctx)

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	var findErr error = nil
	if subgroup != "" {
		findErr = collection.FindOne(context.Background(), bson.M{"name": groupName, "subgroup": subgroup[0]}).Decode(&group)
	} else {
		findErr = collection.FindOne(context.Background(), bson.M{"name": groupName}).Decode(&group)
	}
	if findErr != nil {
		return nil, errors.Wrap(findErr, "Group find error")
	}

	return group, nil
}

func GetGroupList(config *config.MongoConfig, query *config.MongoQuery) (*models.GroupList, error) {
	var result *models.GroupList

	client, ctx, connectErr := connect(config)
	if connectErr != nil {
		return nil, connectErr
	}
	defer disconnect(client, ctx)

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	err := collection.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "Can not find group list")
	}

	return result, nil
}
