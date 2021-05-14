package main

import (
	"github.com/batroff/schedule-back/app"
	"github.com/batroff/schedule-back/app/excel"
	"github.com/batroff/schedule-back/database"
	"github.com/batroff/schedule-back/models/config"
	"github.com/batroff/schedule-back/server"
	"log"
	"os"
)

func main() {
	if os.Args[1] == "--debug" {
		debugMain()
	} else {
		prodMain()
	}
}

func debugMain() {
	//err := app.GetFile("/tmp/hash.xlsx", "https://webservices.mirea.ru/upload/iblock/288/ФТИ_1к_20-21_весна.xlsx")
	//if err != nil {
	//	log.Panic(err)
	//}

	//f, err := os.Open("/tmp/hash.xlsx")
	//if err != nil {
	//	log.Panic(err)
	//}
	//defer f.Close()

	//b, _ := os.ReadFile("/tmp/hash.xlsx")
	//hash = sha1.New()
	//hash.Write

	//hash := sha256.New()
	//if _, err := io.Copy(hash, f); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%x", hash.Sum(nil))
}

func prodMain() {
	cfg := initConfig()

	groups := excel.Parse()

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
