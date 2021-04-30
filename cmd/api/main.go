package main

import (
	"github.com/batroff/schedule-back/app/excel"
	"github.com/batroff/schedule-back/database"
	"github.com/batroff/schedule-back/server"
	"log"
)

func main() {
	groups := excel.Parse()

	err := database.InsertMany("test_database", "test_collection", &groups)

	if err != nil {
		log.Panicf("%v", err)
	}
	err = database.InsertGroupList("test_database", "group_list")
	if err != nil {
		log.Panicf("%v", err)
	}

	server.Start()
}
