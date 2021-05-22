package database

import (
	"context"
	"github.com/batroff/schedule-back/models"
	"github.com/batroff/schedule-back/models/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
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

func InsertGroups(config *config.MongoConfig, query *config.MongoQuery, groups []models.Group) error {
	client, ctx, connectErr := connect(config)
	if connectErr != nil {
		return connectErr
	}
	defer disconnect(client, ctx)

	document := client.Database(query.DocumentName)
	collection := document.Collection(query.CollectionName)

	data := make([]interface{}, len(groups))
	for i, group := range groups {
		data[i] = group
	}

	_, err := collection.InsertMany(ctx, data)
	if err != nil {
		return errors.Wrap(err, "Can not insert groups in db")
	}
	log.Println("Groups successfully uploaded")

	return nil
}

func InsertHash(config *config.MongoConfig, query *config.MongoQuery, hash []string) error {
	client, ctx, connectErr := connect(config)
	if connectErr != nil {
		return connectErr
	}
	defer disconnect(client, ctx)

	collection := client.Database(query.DocumentName).Collection(query.CollectionName)

	_, insertErr := collection.InsertOne(ctx, bson.M{"hash": hash})
	if insertErr != nil {
		return insertErr
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

	collection := client.Database(query.DocumentName).Collection(query.CollectionName)

	var findErr error = nil

	var filter bson.M
	if subgroup != "" {
		subgroupsInt, err := strconv.Atoi(subgroup)
		if err != nil {
			return nil, errors.Wrap(err, "Converting subgroupsInt from string to int error")
		}
		filter = bson.M{"name": groupName, "subgroup": subgroupsInt}
	} else {
		filter = bson.M{"name": groupName}
	}
	findErr = collection.FindOne(context.Background(), filter).Decode(&group)

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

func GetHash(config *config.MongoConfig, query *config.MongoQuery) ([]string, error) {
	var result map[string]interface{}

	client, ctx, connectErr := connect(config)
	if connectErr != nil {
		return nil, connectErr
	}
	defer disconnect(client, ctx)

	collection := client.Database(query.DocumentName).Collection(query.CollectionName)

	err := collection.FindOne(context.Background(), bson.M{}).Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "Can not find hash list")
	}

	var hash []string
	for k, v := range result {
		if k == "hash" {
			for _, h := range v.(primitive.A) {
				hash = append(hash, h.(string))
			}
		}
	}
	return hash, nil
}
