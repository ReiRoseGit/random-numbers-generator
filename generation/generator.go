package generation

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// bound, flows - верхняя граница генерации, количество потоков соответственно
// sortedNumbers - срез отсортированных чисел
// numbersInfo - информация, которая будет защищена от изменения в момент проверки фильтром
type Generator struct {
	sortedNumbers []int
	numbersInfo   information
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
*/
func (g *Generator) filterRandomNumber(channel chan int) {
	number := <-channel
	g.numbersInfo.mutex.Lock()
	if _, ok := g.numbersInfo.usedNumbers[number]; !ok {
		g.numbersInfo.usedNumbers[number] = true
		g.numbersInfo.result = append(g.numbersInfo.result, number)
	}
	g.numbersInfo.mutex.Unlock()
}

// Вызывает несколько (параметр flows) горутин генерации случайного числа (getRandomNumber) и его фильтрации (filterRandomNumber)
func (g *Generator) callAllRoutines(flows int, bound int, channels []chan int) {
	for i := 0; i < flows; i++ {
		go g.getRandomNumber(bound, channels[i])
		go g.filterRandomNumber(channels[i])
	}
}

/*
Функция-диспетчер:
  - Создает срез каналов, по которым передаются данные между
    генератором случайного числа (getRandomNumber) и фильтратором (filterRandomNumber)
  - Запускает функцию callAllRoutines до тех пор, пока количество использованных чисел (usedNumbers) не станет равно
    количеству, равному верхней границе генерации (bound)
*/
func (g *Generator) getAllRandomNumbers(bound, flows int) {
	rand.Seed(time.Now().UnixNano())
	channels := []chan int{}
	for i := 0; i < flows; i++ {
		channels = append(channels, make(chan int))
	}
	for {
		if len(g.numbersInfo.usedNumbers) == bound {
			break
		}
		g.callAllRoutines(flows, bound, channels)
	}
}

// Функция для инициализации. Задает поля и вызывает функцию генерации среза случайных чисел (getAllRandomNumbers)
// так же фиксирует время выполнения генерации среза случайных чисел
func (g *Generator) Generate(bound, flows int) {
	g.numbersInfo.usedNumbers = make(map[int]bool)
	g.numbersInfo.result = []int{}
	start := time.Now()
	g.getAllRandomNumbers(bound, flows)
	duration := time.Since(start)
	fmt.Println("Неотсортированные данные:")
	g.showUnsortedNumbers()
	fmt.Println("Отсортированные данные:")
	g.showSortedNumbers()
	fmt.Println("Время генерации: ", duration)
}

// Выводит в консоль срез случайных чисел
func (g *Generator) showUnsortedNumbers() {
	for _, value := range g.numbersInfo.result {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

// Вызывает функцию (getAndSortNumbers) получения и сортировки нового среза случайных чисел
// И выводит этот срез в консоль
func (g *Generator) showSortedNumbers() {
	g.getAndSortNumbers()
	for _, value := range g.sortedNumbers {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

// Копирует срез случайных чисел, затем сортирует его и записывает в поле sortedNumbers
func (g *Generator) getAndSortNumbers() {
	g.sortedNumbers = make([]int, len(g.numbersInfo.result))
	copy(g.sortedNumbers, g.numbersInfo.result)
	sort.Ints(g.sortedNumbers)
}

// Возвращает структуру с одним публичным методом (Generate)
func NewGenerator() Generator {
	return Generator{}
}
