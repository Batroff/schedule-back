package main

import (
	"fmt"
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
	links, htmlParseErr := html.GetExcelLinks()
	if htmlParseErr != nil {
		log.Panicf("Error occured while html parsing. %v", htmlParseErr)
	}

	excelFiles, downloadErr := app.GetScheduleXlsx("/tmp", links[0])
	if downloadErr != nil {
		log.Panicf("Excel files download error. %v", downloadErr)
	}

	_, parseErr := excel.ParseMultiple(excelFiles)
	if parseErr != nil {
		log.Panicf("%v", parseErr)
	}

	hashes, hashErr := hash.ExcelManyTransform(excelFiles)
	if hashErr != nil {
		log.Panicf("%v", hashErr)
	}
	fmt.Println(hashes)
}

func prodMain() {
	cfg := initConfig()

	links, htmlParseErr := html.GetExcelLinks()
	if htmlParseErr != nil {
		log.Panicf("Error occured while html parsing. %v", htmlParseErr)
	}

	excelFiles, downloadErr := app.GetScheduleXlsx("/tmp", links[0])
	if downloadErr != nil {
		log.Panicf("Excel files download error. %v", downloadErr)
	}

	groups, parseErr := excel.ParseMultiple(excelFiles)
	if parseErr != nil {
		log.Panicf("%v", parseErr)
	}

	dbConfig := database.SetConfig(cfg.Mongo.Host)
	groupsQuery := database.SetQuery("schedule", "groups")
	err := database.InsertGroups(dbConfig, groupsQuery, groups)
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
