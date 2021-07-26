package routes

import (
	"bloms/handlers"
	"net/http"
	"strings"
)

func Routes(w http.ResponseWriter, r *http.Request) {
	var handler func(http.ResponseWriter, *http.Request)
	path := strings.Split(r.URL.Path, "/")[1:]
	pathLength := len(path)
	if pathLength >= 1 {
		switch {
		case path[0] == "":
			handler = handlers.Greet
		case path[0] == "health":
			handler = handlers.HealthCheck
		case path[0] == "message":
			if pathLength == 2 {
				handler = handlers.MessageHandler(path[1])
			} else {
				handler = http.NotFound
			}
		default:
			handler = http.NotFound
		}
	} else {
		handler = http.NotFound
	}
	handler(w, r)
}
