package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/batroff/schedule-back/app"
	"github.com/batroff/schedule-back/database"
	"github.com/batroff/schedule-back/models"
	"log"
	"net/http"
	"regexp"
)

var groupRegexp = regexp.MustCompile(`^[А-Я]{4}[-]\d{2}[-]\d{2}$`)

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	group := query.Get("name")
	subgroup := query.Get("subgroup")

	var response *models.ResponseGroup

	if groupRegexp.MatchString(group) {
		cfg, err := app.LoadConfig("config/")
		if err != nil {
			log.Panicf("%s", err)
		}

		groupsQuery := database.SetQuery("schedule", "groups")
		dbGroup, err := database.FindGroup(cfg.Mongo, groupsQuery, group, subgroup)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			response = &models.ResponseGroup{
				ErrorMsg: "no documents in result",
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
