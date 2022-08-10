package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"random-numbers-generator/generation"
)

// generator - объект-генератор
// numberInfo - информация о текущем сгенерированном срезе
type numberGenerator struct {
	generator  generation.Generator
	numberInfo NumbersInformation
}

// Структура, которая в последующем будет преобразовываться в JSON
type NumbersInformation struct {
	UnsortedNumbers []int         `json:"unsorted_numbers"`
	SortedNumbers   []int         `json:"sorted_numbers"`
	Time            time.Duration `json:"time"`
}

// Конструктор генератора, вызывается один раз в пакете main
func NewNumberGenerator() numberGenerator {
	return numberGenerator{generator: generation.NewGenerator(), numberInfo: NumbersInformation{}}
}

// Генерирует JSON файл
func (ng *numberGenerator) getJSON(w http.ResponseWriter, r *http.Request, bound, flows int) {
	var js []byte
	if bound < 1 || flows < 1 {
		js, _ = json.Marshal(&NumbersInformation{UnsortedNumbers: []int{}, SortedNumbers: []int{}})
	} else {
		unsortedNumbers, sortedNumbers, time := ng.generator.Generate(bound, flows)
		js, _ = json.Marshal(&NumbersInformation{UnsortedNumbers: unsortedNumbers, SortedNumbers: sortedNumbers, Time: time})
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Получает request и значение, которое нужно найти в качестве query-параметра
// Конвертирует в int найденный query-параметр
func (ng *numberGenerator) getQuery(r *http.Request, q string) (int, error) {
	return strconv.Atoi(r.URL.Query().Get(q))
}

// Парсит query параметры и вызывает функцию для генерации JSON с ними
func (ng *numberGenerator) getQueriesAndJSON(w http.ResponseWriter, r *http.Request) {
	bound, err1 := ng.getQuery(r, "bound")
	flows, err2 := ng.getQuery(r, "flows")
	if err1 != nil || err2 != nil {
		http.NotFound(w, r)
		return
	}
	ng.getJSON(w, r, bound, flows)
}

// Обрабатывает маршрут /numbers
func (ng *numberGenerator) NumbersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ng.getQueriesAndJSON(w, r)
	} else {
		http.Error(w, fmt.Sprintf("expect method GET /numbers, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}
