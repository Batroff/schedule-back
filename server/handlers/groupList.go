package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/batroff/schedule-back/app"
	"github.com/batroff/schedule-back/database"
	"github.com/batroff/schedule-back/models"
	"log"
	"net/http"
)

func GroupListHandler(w http.ResponseWriter, _ *http.Request) {
	var response *models.ResponseGroupList

	cfg, err := app.LoadConfig("./config/app.yaml")
	if err != nil {
		log.Panicf("%s", err)
	}

	groupsQuery := database.SetQuery("schedule", "group_list")
	groupList, err := database.GetGroupList(cfg.Mongo, groupsQuery)
	if err == nil {
		response = &models.ResponseGroupList{
			GroupList: groupList.Map,
			ErrorMsg:  "",
		}
		w.WriteHeader(http.StatusOK)
	} else {
		log.Printf("Group list error -- %s", err)
		response = &models.ResponseGroupList{
			GroupList: nil,
			ErrorMsg:  "group list is empty",
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}
