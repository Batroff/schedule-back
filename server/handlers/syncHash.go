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

func HashHandler(w http.ResponseWriter, _ *http.Request) {
	var response *models.ResponseHash

	cfg, err := app.LoadConfig("./config/app.yaml")
	if err != nil {
		log.Panicf("Hash handler error. %v", err)
	}

	hashQuery := database.SetQuery("schedule", "hash")
	hash, err := database.GetHash(cfg.Mongo, hashQuery)

	if err == nil {
		response = &models.ResponseHash{
			Hash:     hash,
			ErrorMsg: "",
		}
		w.WriteHeader(http.StatusOK)
	} else {
		log.Printf("Hash error -- %s", err)
		response = &models.ResponseHash{
			Hash:     nil,
			ErrorMsg: "hash list is empty",
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}
