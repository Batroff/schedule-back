package server

import (
	"github.com/batroff/schedule-back/app"
	"github.com/batroff/schedule-back/server/handlers"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/api/groupList/", handlers.GroupListHandler)
	http.HandleFunc("/api/group/", handlers.GroupHandler)
	http.HandleFunc("/api/hash/", handlers.HashHandler)

	cfg, err := app.LoadConfig("config/")
	if err != nil {
		log.Panicf("%s", err)
	}
	log.Printf("Запуск сервера на %s...\n", cfg.Server.Address)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
