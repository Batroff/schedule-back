package server

import (
	"encoding/json"
	"fmt"
	"github.com/batroff/schedule-back/database"
	"github.com/batroff/schedule-back/models"
	"log"
	"net/http"
	"regexp"
)

func Start() {
	http.HandleFunc("/api/groupList/", groupListHandler)
	http.HandleFunc("/api/group/", find)

	log.Println("Запуск сервера на 127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var groupRegexp = regexp.MustCompile(`^[А-Я]{4}[-]\d{2}[-]\d{2}$`)

func find(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	group := query.Get("name")
	subgroup := query.Get("subgroup")

	var response *models.ResponseGroup

	if groupRegexp.MatchString(group) {
		dbGroup, err := database.FindGroup("test_database", "test_collection", group, subgroup)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			response = &models.ResponseGroup{
				ErrorMsg: err.Error(),
				Group:    dbGroup,
			}
		} else {
			w.WriteHeader(http.StatusOK)
			response = &models.ResponseGroup{
				ErrorMsg: "",
				Group:    dbGroup,
			}
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		response = &models.ResponseGroup{
			ErrorMsg: "Bad request",
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

func groupListHandler(w http.ResponseWriter, r *http.Request) {
	var response *models.ResponseGroupList

	groupList, err := database.GetGroupList("test_database", "group_list")
	if err == nil {
		response = &models.ResponseGroupList{
			GroupList: groupList,
			ErrorMsg:  "",
		}
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Println(err)
		response = &models.ResponseGroupList{
			GroupList: models.GroupList{},
			ErrorMsg:  "GroupList is empty",
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}
