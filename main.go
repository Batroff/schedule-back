package main

import (
	"log"
	"schedule/db"
	"schedule/parse"
)

func main() {
	groups, groupsMini := parse.Parse()
	log.Printf(string(len(groups)))
	log.Printf(string(len(groupsMini)))
	//for _, mini := range groupsMini {
	//	fmt.Println(mini)
	//}
	err := db.InsertMany("test_database", "test_collection", &groups /*Mini*/)

	if err != nil {
		log.Panicf("%v", err)
	}
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
