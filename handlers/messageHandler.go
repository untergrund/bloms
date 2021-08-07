package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"
)

func MessageHandler(path string) func(w http.ResponseWriter, r *http.Request) {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable ot connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO public.messages (message) VALUES($1);", path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}

	log.Println("message/" + path)
	response := "message insert to db: " + path
	respJSON, _ := json.Marshal(response)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	}
}
