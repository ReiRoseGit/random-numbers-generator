package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"random-numbers-generator/generation"
)

type NumbersInformation struct {
	UnsortedNumbers []int
	SortedNumbers   []int
	Time            time.Duration
}

// Обрабатывает маршрут /
// Выводит на экран информацию о получении случайных значений
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "bound - верхяя граница генерации, flows - количество потоков")
	fmt.Fprintln(w, "Для получения случайных чисел перейдите по адресу:")
	fmt.Fprintln(w, "/numbers?bound=<значение>&flows=<значение>")
}

// Получает request и значение, которое нужно найти в качестве query-параметра
// Конвертирует в int найденный query-параметр
func getQuery(r *http.Request, q string) (int, error) {
	return strconv.Atoi(r.URL.Query().Get(q))
}

// Обрабатывает маршрут /numbers
// Генерирует список случайных чисел на основании query-параметров: bound и flows
func generateNumbers(w http.ResponseWriter, r *http.Request) {
	bound, err1 := getQuery(r, "bound")
	flows, err2 := getQuery(r, "flows")
	if err1 != nil || bound < 1 || err2 != nil || flows < 1 {
		http.NotFound(w, r)
		return
	}
	g := generation.NewGenerator()
	jsonNumbers, jsonErr := createJSON(g.Generate(bound, flows))
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Fprintln(w, string(jsonNumbers))
}

// Генерирует и возвращает JSON
func createJSON(unsortedNumbers []int, sortedNumbers []int, time time.Duration) ([]byte, error) {
	return json.Marshal(&NumbersInformation{UnsortedNumbers: unsortedNumbers, SortedNumbers: sortedNumbers, Time: time})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/numbers", generateNumbers)
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
