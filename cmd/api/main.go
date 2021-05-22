package main

import (
	"github.com/batroff/schedule-back/app"
	"github.com/batroff/schedule-back/app/excel"
	"github.com/batroff/schedule-back/app/hash"
	"github.com/batroff/schedule-back/app/html"
	"github.com/batroff/schedule-back/database"
	"github.com/batroff/schedule-back/models/config"
	"github.com/batroff/schedule-back/server"
	"log"
	"os"
)

func main() {
	for _, arg := range os.Args {
		if arg == "--debug" {
			debugMain()
			return
		}
	}

	prodMain()
}

func debugMain() {
	server.Start()
}

func prodMain() {
	cfg := initConfig()

	// Download html mirea.ru/schedule/
	// Parse excel links from html file
	links, htmlParseErr := html.GetExcelLinks()
	if htmlParseErr != nil {
		log.Panicf("Error occured while html parsing. %v", htmlParseErr)
	}

	// Download all excel files to ./tmp/schedule_#.xlsx
	excelFiles, downloadErr := app.GetScheduleXlsx("/tmp", links[0])
	if downloadErr != nil {
		log.Panicf("Excel files download error. %v", downloadErr)
	}

	// Parse all excel files
	groups, parseErr := excel.ParseMultiple(excelFiles)
	if parseErr != nil {
		log.Panicf("%v", parseErr)
	}

	// Transform all excel files to hash array
	hashes, hashErr := hash.ExcelManyTransform(excelFiles)
	if hashErr != nil {
		log.Panicf("%v", hashErr)
	}

	// Set database config
	dbConfig := database.SetConfig(cfg.Mongo.Host)

	// Add groups to database
	groupsQuery := database.SetQuery("schedule", "groups")
	err := database.InsertGroups(dbConfig, groupsQuery, groups)
	if err != nil {
		log.Panicf("%v", err)
	}

	// Add list of all groups to database
	groupListQuery := database.SetQuery("schedule", "group_list")
	err = database.InsertGroupList(dbConfig, groupListQuery)
	if err != nil {
		log.Panicf("%v", err)
	}

	// Add hash of excel files to database
	hashQuery := database.SetQuery("schedule", "hash")
	hashErr = database.InsertHash(dbConfig, hashQuery, hashes)
	if hashErr != nil {
		log.Panicf("%v", hashErr)
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
