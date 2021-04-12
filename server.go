package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"schedule/db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	group := query.Get("group")

	dbGroup := db.FindGroup("test_database", "test_collection", group)

	response := dbGroup

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}
