package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	respJSON, _ := json.Marshal(currentTime)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	response := "Hello from Bloms!"
	respJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}
