package main

import (
	"log"
	"net/http"
	"random-numbers-generator/routing"
)

func main() {
	mux := http.NewServeMux()
	generator := routing.NewNumberGenerator()
	mux.HandleFunc("/ws", generator.WebSocketHandler)
	mux.Handle("/", http.FileServer(http.Dir("dist")))
	mux.HandleFunc("/numbers", generator.NumbersHandler)
	mux.HandleFunc("/history", generator.HistoryHandler)
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
