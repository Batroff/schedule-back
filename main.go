package main

import (
	"log"
	"schedule/db"
	"schedule/parse"
)

func main() {
	groups := parse.Parse()

	err := db.InsertMany("test_database", "test_collection", &groups)

	if err != nil {
		log.Panicf("%v", err)
	}
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
