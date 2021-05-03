package main

import (
	"github.com/batroff/schedule-back/app"
	"github.com/batroff/schedule-back/app/excel"
	"github.com/batroff/schedule-back/database"
	"github.com/batroff/schedule-back/models/config"
	"github.com/batroff/schedule-back/server"
	"log"
)

func main() {
	cfg := initConfig()

	groups := excel.Parse()

	dbConfig := database.SetConfig(cfg.Mongo.Host)
	groupsQuery := database.SetQuery("schedule", "groups")
	err := database.InsertGroups(dbConfig, groupsQuery, &groups)
	if err != nil {
		log.Panicf("%v", err)
	}

	groupListQuery := database.SetQuery("schedule", "group_list")
	err = database.InsertGroupList(dbConfig, groupListQuery)
	if err != nil {
		log.Panicf("%v", err)
	}

	server.Start()
}

func initConfig() *config.AppConfig {
	cfg, err := app.LoadConfig("config/")
	if err != nil {
		log.Panicf("%s", err)
	}

	return cfg
}
