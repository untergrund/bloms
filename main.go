package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := "Hello fron Bloms!"
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})
	port := ":8080"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}
