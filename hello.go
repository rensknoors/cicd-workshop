package main

import (
	"HelloGo/Handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", Handlers.HealthCheckHandler)
	http.HandleFunc("/", Handlers.TalkBackHandler)
	log.Println("Starting server");
	http.ListenAndServe(":8080", nil)
}

