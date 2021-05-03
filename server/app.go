package server

import (
	"github.com/batroff/schedule-back/server/handlers"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/api/groupList/", handlers.GroupListHandler)
	http.HandleFunc("/api/group/", handlers.GroupHandler)

	log.Println("Запуск сервера на 127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
