package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"random-numbers-generator/generation"
)

type NumbersInformation struct {
	UnsortedNumbers []int         `json:"unsorted_numbers"`
	SortedNumbers   []int         `json:"sorted_numbers"`
	Time            time.Duration `json:"time"`
}

// Получает request и значение, которое нужно найти в качестве query-параметра
// Конвертирует в int найденный query-параметр
func getQuery(r *http.Request, q string) (int, error) {
	return strconv.Atoi(r.URL.Query().Get(q))
}

// Обрабатывает маршрут /numbers
// Генерирует список случайных чисел на основании query-параметров: bound и flows
func numbersHandler(w http.ResponseWriter, r *http.Request) {
	bound, err1 := getQuery(r, "bound")
	flows, err2 := getQuery(r, "flows")
	if err1 != nil || err2 != nil {
		http.NotFound(w, r)
		return
	}
	g := generation.NewGenerator()
	renderJSON(w, &g, bound, flows)
}

// Генерирует JSON
func renderJSON(w http.ResponseWriter, g *generation.Generator, bound, flows int) {
	var js []byte
	if bound < 1 || flows < 1 {
		js, _ = json.Marshal(&NumbersInformation{UnsortedNumbers: []int{}, SortedNumbers: []int{}})
	} else {
		unsortedNumbers, sortedNumbers, time := g.Generate(bound, flows)
		js, _ = json.Marshal(&NumbersInformation{UnsortedNumbers: unsortedNumbers, SortedNumbers: sortedNumbers, Time: time})
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/numbers", numbersHandler)
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
