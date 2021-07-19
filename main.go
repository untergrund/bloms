package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type response struct {
	Message string `json:"message"`
}

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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Greet)
	mux.HandleFunc("/health", HealthCheck)

	port := ""
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8000"
	}

	server := &http.Server{
		Handler:        mux,
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
