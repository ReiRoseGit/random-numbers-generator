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
	bound, flows  int
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
func (g *Generator) getRandomNumber(channel chan int) {
	channel <- 1 + rand.Intn(g.bound)
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
func (g *Generator) callAllRoutines(channels []chan int) {
	for i := 0; i < g.flows; i++ {
		go g.getRandomNumber(channels[i])
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
func (g *Generator) getAllRandomNumbers() {
	rand.Seed(time.Now().UnixNano())
	channels := []chan int{}
	for i := 0; i < g.flows; i++ {
		channels = append(channels, make(chan int))
	}
	for {
		if len(g.numbersInfo.usedNumbers) == g.bound {
			break
		}
		g.callAllRoutines(channels)
	}
}

// Функция для инициализации. Задает поля и вызывает функцию генерации среза случайных чисел (getAllRandomNumbers)
// так же фиксирует время выполнения генерации среза случайных чисел
func (g *Generator) Generate() {
	g.numbersInfo.usedNumbers = make(map[int]bool)
	g.numbersInfo.result = []int{}
	start := time.Now()
	g.getAllRandomNumbers()
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
func NewGenerator(bound, flows int) Generator {
	return Generator{bound: bound, flows: flows}
}
