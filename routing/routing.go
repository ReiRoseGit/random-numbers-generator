package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"random-numbers-generator/basing"
	"random-numbers-generator/generation"

	"github.com/gorilla/websocket"
)

// generator - объект-генератор
// numberInfo - информация о текущем сгенерированном срезе
// db - база данных в виде json
type numberGenerator struct {
	generator  generation.Generator
	numberInfo NumbersInformation
	db         basing.Driver
}

// Структура, которая в последующем будет преобразовываться в JSON
type NumbersInformation struct {
	UnsortedNumbers []int         `json:"unsorted_numbers"`
	SortedNumbers   []int         `json:"sorted_numbers"`
	Time            time.Duration `json:"time"`
}

// Структура, описывающая параметры
type Params struct {
	Bound string `json:"bound"`
	Flows string `json:"flows"`
}

// Конструктор генератора, вызывается один раз в пакете main
func NewNumberGenerator() numberGenerator {
	return numberGenerator{generator: generation.NewGenerator(), numberInfo: NumbersInformation{}, db: *basing.New("db.json")}
}

// Статически генерирует JSON файл, записывает последнюю генерацию в БД
func (ng *numberGenerator) writeHttpJSON(ctx context.Context, w http.ResponseWriter, r *http.Request, bound, flows int) {
	var js []byte
	unsortedChannel := make(chan []int)
	timeChannel := make(chan time.Duration)
	go ng.generator.Generate(ctx, unsortedChannel, timeChannel, bound, flows)
	unsortedSlice := <-unsortedChannel
	generationTime := <-timeChannel
	sortedSlice := sortSlice(unsortedSlice)
	ng.db.AddJSON(unsortedSlice, sortedSlice, generationTime, time.Now().Format("01-02-2006 15:04:05"))
	js, _ = json.Marshal(&NumbersInformation{UnsortedNumbers: unsortedSlice, SortedNumbers: sortedSlice, Time: generationTime})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Обрабатывает маршрут /numbers
func (ng *numberGenerator) NumbersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		bound, _ := strconv.Atoi(r.FormValue("bound"))
		flows, _ := strconv.Atoi(r.FormValue("flows"))
		ng.writeHttpJSON(ctx, w, r, bound, flows)
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
// генерирующей последовательность, так же создает контекст, отменяет его в случае потери соединения с клиентом
func (ng *numberGenerator) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		mt, p, err := connection.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			cancel()
			break
		}
		var params Params
		json.Unmarshal(p, &params)
		bound, _ := strconv.Atoi(params.Bound)
		flows, _ := strconv.Atoi(params.Flows)
		go ng.liveNumbers(ctx, connection, bound, flows)
	}
}

// Создает каналы для обмена информацией: динамический канал, канал неотсортированных чисел и времени, динамически отправляет числа клиенту
// до тех пор, пока либо не сгенерирует нужное количество, либо не будет потеряно соединение с клиентом.
// В случае успешной генерации запускает функцию отправки JSON клиенту, в конце производится очистка каналов
func (ng *numberGenerator) liveNumbers(ctx context.Context, connection *websocket.Conn, bound, flows int) {
	liveChannel := make(chan int)
	unsortedChannel := make(chan []int)
	timeChannel := make(chan time.Duration)
	go ng.generator.Generate(ctx, unsortedChannel, timeChannel, bound, flows, liveChannel)
	var unsorted []int
loop:
	for {
		select {
		// Динамическая передача числа на вывод
		case value := <-liveChannel:
			connection.WriteMessage(1, []byte(strconv.Itoa(value)))
		// Передача готового среза чисел
		case value := <-unsortedChannel:
			unsorted = value
			break loop
		case <-ctx.Done():
			break loop
		}
	}
	// Если соединение не было прервано
	if len(unsorted) != 0 {
		ng.writeWebsocketJSON(connection, <-timeChannel, unsorted)
	}
	ng.clearChannels(liveChannel, unsortedChannel)
}

// Очистка канала, по которому передаются сгенерированные числа
func (ng *numberGenerator) clearChannels(liveChannel chan int, unsortedChannel chan []int) {
	for {
		if _, ok := <-liveChannel; !ok {
			break
		}
	}
}

// Сортирует срез, переданный ему в качестве параметра и возвращает готовый результат
func sortSlice(unsorted []int) []int {
	sorted := make([]int, len(unsorted))
	copy(sorted, unsorted)
	sort.Ints(sorted)
	return sorted
}

// Вызывает функцию получения отсортированного среза и записывает все данные в JSON, затем отправляет его клиенту, записывает последнюю генерацию в БД
func (ng *numberGenerator) writeWebsocketJSON(connection *websocket.Conn, t time.Duration, unsorted []int) {
	sorted := sortSlice(unsorted)
	data := NumbersInformation{UnsortedNumbers: unsorted, SortedNumbers: sorted, Time: t}
	ng.db.AddJSON(unsorted, sorted, t, time.Now().Format("01-02-2006 15:04:05"))
	js, _ := json.Marshal(&data)
	connection.WriteMessage(1, js)
}
