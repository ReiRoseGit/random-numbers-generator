package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"random-numbers-generator/generation"
)

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
	unsortedNumbers, sortedNumbers, time := g.Generate(bound, flows)
	showNumbersInfo(w, unsortedNumbers, sortedNumbers, time)
}

// Вывод данных на экран
func showNumbersInfo(w http.ResponseWriter, unsortedNumbers []int, sortedNumbers []int, time time.Duration) {
	fmt.Fprintln(w, "Неотсортированные данные:")
	fmt.Fprintln(w, unsortedNumbers)
	fmt.Fprintln(w, "Отсортированные данные:")
	fmt.Fprintln(w, sortedNumbers)
	fmt.Fprintln(w, "Время:")
	fmt.Fprintln(w, time)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/numbers", generateNumbers)
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
