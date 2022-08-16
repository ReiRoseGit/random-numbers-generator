package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"random-numbers-generator/generation"

	"github.com/gorilla/websocket"
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

// Структура, описывающая параметры
type Params struct {
	Bound string `json:"bound"`
	Flows string `json:"flows"`
}

// Конструктор генератора, вызывается один раз в пакете main
func NewNumberGenerator() numberGenerator {
	return numberGenerator{generator: generation.NewGenerator(), numberInfo: NumbersInformation{}}
}

// Генерирует JSON файл
func (ng *numberGenerator) getJSON(w http.ResponseWriter, r *http.Request, bound, flows int) {
	var js []byte
	unsortedNumbers, sortedNumbers, time := ng.generator.Generate(bound, flows)
	js, _ = json.Marshal(&NumbersInformation{UnsortedNumbers: unsortedNumbers, SortedNumbers: sortedNumbers, Time: time})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Запускает функцию генерации последовательностей и предает сгенерированные данные в соответствующие каналы
func (ng *numberGenerator) getLiveNumbers(bound, flows int,
	liveChannel chan int, sortedChannel chan []int, timeChannel chan time.Duration, unsortedChannel chan []int) {
	unsortedNumbers, sortedNumbers, time := ng.generator.Generate(bound, flows, liveChannel)
	unsortedChannel <- unsortedNumbers
	sortedChannel <- sortedNumbers
	timeChannel <- time
}

// Обрабатывает маршрут /numbers
func (ng *numberGenerator) NumbersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bound, _ := strconv.Atoi(r.FormValue("bound"))
		flows, _ := strconv.Atoi(r.FormValue("flows"))
		ng.getJSON(w, r, bound, flows)
	} else {
		http.Error(w, fmt.Sprintf("expect method Post, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// Настройка websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Управление websocket. Принимает значения из формы и создает каналы для приема данных из функции getLiveNumbers,
// генерирующей последовательность
func (ng *numberGenerator) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()
	for {
		mt, p, err := connection.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}
		var params Params
		json.Unmarshal(p, &params)
		bound, _ := strconv.Atoi(params.Bound)
		flows, _ := strconv.Atoi(params.Flows)
		ng.liveNumbers(connection, bound, flows)
	}
}

// Создает каналы для обмена информацией, динамически отправляет числа клиенту,
// формирует и отправляет JSON файл
func (ng *numberGenerator) liveNumbers(connection *websocket.Conn, bound, flows int) {
	// Канал для динамического вывода чисел
	liveChannel := make(chan int)
	unsortedChannel := make(chan []int)
	sortedChannel := make(chan []int)
	timeChannel := make(chan time.Duration)
	go ng.getLiveNumbers(bound, flows, liveChannel, sortedChannel, timeChannel, unsortedChannel)
	var sorted, unsorted []int
	var time time.Duration
	for {
		if bound == 0 {
			break
		}
		select {
		case value := <-liveChannel:
			connection.WriteMessage(1, []byte(strconv.Itoa(value)))
		case value := <-unsortedChannel:
			unsorted = value
		case value := <-sortedChannel:
			sorted = value
		case value := <-timeChannel:
			time = value
			bound = 0
		}
	}
	js, _ := json.Marshal(&NumbersInformation{UnsortedNumbers: unsorted, SortedNumbers: sorted, Time: time})
	connection.WriteMessage(1, js)
}
