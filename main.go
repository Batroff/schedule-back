package main

import (
	"log"
	"net/http"
	"schedule/db"
	"schedule/parse"
)

func main() {
	groups := parse.Parse()

	err := db.InsertMany("test_database", "test_collection", &groups)

	if err != nil {
		log.Panicf("%v", err)
	}
	//os.Getenv()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/group", find)
	log.Println("Запуск сервера на 127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
