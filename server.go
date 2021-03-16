package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	group := query.Get("group")

	response := HelloResponse{
		Message: fmt.Sprintf("Group: %s", group),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// 1st way
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}

	// 2nd way
	//jData, err := json.Marshal(response)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//w.Write(jData)
}
