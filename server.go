package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"schedule/db"
	"schedule/structure"
)

type ResponseGroup struct {
	ErrorMsg string          `json:"ErrorMsg,omitempty"`
	Group    structure.Group `json:"Group,omitempty"`
}

type ResponseGroupList struct {
	ErrorMsg  string              `json:"ErrorMsg,omitempty"`
	GroupList structure.GroupList `json:"GroupList"`
}

var groupRegexp = regexp.MustCompile(`[А-Я]{4}[-]\d{2}[-]\d{2}`)

func find(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	group := query.Get("group")
	subgroup := query.Get("subgroup")

	var response *ResponseGroup

	if groupRegexp.MatchString(group) {
		dbGroup, err := db.FindGroup("test_database", "test_collection", group, subgroup)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			response = &ResponseGroup{
				ErrorMsg: err.Error(),
				Group:    dbGroup,
			}
		} else {
			w.WriteHeader(http.StatusOK)
			response = &ResponseGroup{
				ErrorMsg: "",
				Group:    dbGroup,
			}
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		response = &ResponseGroup{
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
	var response *ResponseGroupList

	groupList, err := db.GetGroupList("test_database", "group_list")
	if err == nil {
		response = &ResponseGroupList{
			GroupList: groupList,
			ErrorMsg:  "",
		}
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Println(err)
		response = &ResponseGroupList{
			GroupList: structure.GroupList{},
			ErrorMsg:  "GroupList is empty",
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}
