package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func MessageHandler(path string) func(w http.ResponseWriter, r *http.Request) {
	log.Println("message/" + path)
	response := "message/" + path + " endpoint"
	respJSON, _ := json.Marshal(response)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	}
}
