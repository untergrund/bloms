package main

import (
	"bloms/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ""
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8000"
	}

	server := &http.Server{
		Handler:        http.HandlerFunc(routes.Routes),
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
