package main

import (
	"log"
	"net/http"
	"random-numbers-generator/routing"
)

func main() {
	mux := http.NewServeMux()
	generator := routing.NewNumberGenerator()
	mux.Handle("/", http.FileServer(http.Dir("dist")))
	mux.HandleFunc("/numbers", generator.NumbersHandler)
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
