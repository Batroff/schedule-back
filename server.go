package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"schedule/db"
	"schedule/structure"
)

// TODO: Rework struct to interface{}
type Response struct {
	ErrorMsg string `json:",omitempty"`
	Group    structure.Group
}

var groupRegexp = regexp.MustCompile(`[А-Я]{4}[-]\d{2}[-]\d{2}`)

func find(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	group := query.Get("group")

	var response *Response

	if groupRegexp.MatchString(group) {
		dbGroup := db.FindGroup("test_database", "test_collection", group)
		response = &Response{
			ErrorMsg: "",
			Group:    dbGroup,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	query := r.URL.Query()
//	group := query.Get("group")
//
//	var response *Response
//
//	if groupRegexp.MatchString(group) {
//		dbGroup := db.FindGroup("test_database", "test_collection", group)
//		response = &Response{
//			ErrorMsg: "",
//			Group:    dbGroup,
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//	} else {
//		w.WriteHeader(http.StatusBadRequest)
//	}
//
//	err := json.NewEncoder(w).Encode(response)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
