package generation

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

// bound, flows - верхняя граница генерации, количество потоков соответственно
// sortedNumbers - срез отсортированных чисел
// numbersInfo - информация, которая будет защищена от изменения в момент проверки фильтром
type Generator struct {
	numbersInfo information
}

// result - срез случайных чисел
// usedNumbers - словарь чисел, которые уже были использованы при генерации
type information struct {
	mutex       sync.Mutex
	result      []int
	usedNumbers map[int]bool
}

// Генерирует и отправляет в канал случайное число из отрезка [1, bound]
func (g *Generator) getRandomNumber(bound int, channel chan int) {
	channel <- 1 + rand.Intn(bound)
}

/*
Фильтрует случайное число, опираясь на словарь usedNumbers, следующим образом:
- Извлечение значения из канала
- Фиксирование значения словаря с помощью мьютекса
- Сравнение полученного по каналу значения со значениями словаря usedNumbers
- Добавления значения в словарь использованных чисел (usedNumbers) и в срез-результат (result) в случае уникальности числа
- В случае динамического отправления данных (liveChannel - не пустой параметр), передает в этот канал уникальное число
*/
func (g *Generator) filterRandomNumber(channel chan int, liveChannel ...chan int) {
	number := <-channel
	g.numbersInfo.mutex.Lock()
	if _, ok := g.numbersInfo.usedNumbers[number]; !ok {
		g.numbersInfo.usedNumbers[number] = true
		g.numbersInfo.result = append(g.numbersInfo.result, number)
		if len(liveChannel) == 1 {
			liveChannel[0] <- number
		}
	}
	g.numbersInfo.mutex.Unlock()
}

// Вызывает несколько (параметр flows) горутин генерации случайного числа (getRandomNumber) и его фильтрации (filterRandomNumber)
func (g *Generator) callAllRoutines(flows int, bound int, channels []chan int, liveChannel ...chan int) {
	for i := 0; i < flows; i++ {
		go g.getRandomNumber(bound, channels[i])
		go g.filterRandomNumber(channels[i], liveChannel...)
	}
}

// Создает каналы для генерации и фильтрации случайного числа, запускает горутины до тех пор, пока не будет прервано соединение с клиентом
// или  пока не сгенерирует все необходимые числа, фиксирует время генерации всех чисел и отправляет в канал времени
func (g *Generator) getAllRandomNumbers(ctx context.Context, bound, flows int, result chan []int, timeCh chan time.Duration, channel ...chan int) {
	rand.Seed(time.Now().UnixNano())
	channels := []chan int{}
	// Каналы для генерации и фильтрации
	for i := 0; i < flows; i++ {
		channels = append(channels, make(chan int))
	}
	// Если процесс генерации прерван
	killed := ctx.Done()
	// Засекает время начала генерации
	start := time.Now()
loop:
	for {
		select {
		// Соединение прервано
		case <-killed:
			result <- []int{}
			timeCh <- 0
			break loop
		default:
		}
		// Все числа сгенерированы
		if len(g.numbersInfo.usedNumbers) == bound {
			result <- g.numbersInfo.result
			timeCh <- time.Since(start)
			break loop
		}
		// Вызов необходимого числа горутин
		g.callAllRoutines(flows, bound, channels, channel...)
	}
}

// Функция для инициализации. Задает поля и вызывает функцию генерации среза случайных чисел (getAllRandomNumbers)
func (g *Generator) Generate(ctx context.Context, result chan []int, timeCh chan time.Duration, bound, flows int, channel ...chan int) {
	g.numbersInfo.usedNumbers = make(map[int]bool)
	g.numbersInfo.result = []int{}
	go g.getAllRandomNumbers(ctx, bound, flows, result, timeCh, channel...)
}

// Возвращает структуру с одним публичным методом (Generate)
func NewGenerator() Generator {
	return Generator{}
}
