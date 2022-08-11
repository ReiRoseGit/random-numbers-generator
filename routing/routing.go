package routing

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

// Структура, описывающая JSON файл при некорректных данных
type ErrorJSON struct {
	ErrCode    int    `json:"error_code"`
	ErrMessage string `json:"error_message"`
}

type htmlData struct {
	Title string
}

// Конструктор генератора, вызывается один раз в пакете main
func NewNumberGenerator() numberGenerator {
	return numberGenerator{generator: generation.NewGenerator(), numberInfo: NumbersInformation{}}
}

// Генерирует JSON файл
func (ng *numberGenerator) getJSON(w http.ResponseWriter, r *http.Request, bound, flows int) {
	var js []byte
	code := ng.cleanParams(bound, flows)
	switch code {
	case 0:
		unsortedNumbers, sortedNumbers, time := ng.generator.Generate(bound, flows)
		js, _ = json.Marshal(&NumbersInformation{UnsortedNumbers: unsortedNumbers, SortedNumbers: sortedNumbers, Time: time})
	case 100:
		js, _ = json.Marshal(&ErrorJSON{ErrCode: code, ErrMessage: "bound parameter must be a positive number"})
	case 200:
		js, _ = json.Marshal(&ErrorJSON{ErrCode: code, ErrMessage: "flows parameter must be a positive number"})
	case 300:
		js, _ = json.Marshal(&ErrorJSON{ErrCode: code, ErrMessage: "parameters must be positive numbers"})
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Валидатор параметров, возвращает код ошибки или 0, в случае, если данные валидны
func (ng *numberGenerator) cleanParams(bound, flows int) int {
	if bound > 0 && flows > 0 {
		return 0
	} else if bound < 1 && flows < 1 {
		return 300
	} else if bound < 1 {
		return 100
	} else {
		return 200
	}
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
	} else if r.Method == http.MethodPost {
		bound := r.FormValue("bound")
		flows := r.FormValue("flows")
		w.Write([]byte(bound))
		fmt.Fprintln(w)
		w.Write([]byte(flows))
	} else {
		http.Error(w, fmt.Sprintf("expect method Get, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// Обрабатывает маршрут /
func (ng *numberGenerator) IndexHandler(w http.ResponseWriter, r *http.Request) {
	str, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	tmpl, err := template.ParseFiles(strings.Join([]string{str, "/templates/", "index.html"}, ""))
	if err != nil {
		log.Fatal("Unable to parse from template:", err)
	}
	tmpl.Execute(w, nil)
}
