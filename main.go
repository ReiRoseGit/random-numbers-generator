package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// функция для получения одного случайного числа
func getRandomNumber(count int, channel chan int) {
	channel <- 1 + rand.Intn(count)
}

// функция-фильтр для проверки допустимых чисел, возвращает обновленные: словарь использованных чисел и срез-результат
// flows - количество потоков
// channels - каналы случайных чисел
// usedNumbers - использованные числа
// result - slice, являющийся результатом
func filterRandomNumbers(flows int, channels []chan int, usedNumbers map[int]bool, result []int) ([]int, map[int]bool) {
	for i := 0; i < flows; i++ {
		number := <-channels[i]
		if _, ok := usedNumbers[number]; !ok {
			result = append(result, number)
			usedNumbers[number] = true
		}
	}
	return result, usedNumbers
}

// фнукция-диспетчер
func getAllRandomNumbers(count, flows int) []int {
	rand.Seed(time.Now().UnixNano())
	result := []int{}
	usedNumbers := make(map[int]bool)
	channels := []chan int{}
	// создание каналов
	for i := 0; i < flows; i++ {
		channels = append(channels, make(chan int))
	}
	for {
		if len(result) == count {
			break
		}
		// логика по запуску горутин с каналами
		for i := 0; i < flows; i++ {
			go getRandomNumber(count, channels[i])

		}
		// логика по проверке доступных значений
		result, usedNumbers = filterRandomNumbers(flows, channels, usedNumbers, result)
	}
	return result
}

// вывод чисел на экран
func showRandomNumbers(numbers []int) {
	for _, value := range numbers {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

// сортировка чисел в порядке возрастания
func inplaceSortRandomNumbers(numbers []int) {
	sort.Ints(numbers)
}

func main() {
	// флаги:
	// bound - верхняя граница
	// flows - количество потоков
	var bound, flows int
	flag.IntVar(&bound, "bound", 30, "Верхняя граница (положительное целое число)")
	flag.IntVar(&flows, "flows", 3, "Количество потоков (положительное целое число)")
	flag.Parse()
	if bound < 1 || flows < 1 {
		flag.Usage()
		os.Exit(1)
	} else {
		result := getAllRandomNumbers(bound, flows)
		fmt.Println("Неотсортированные данные:")
		showRandomNumbers(result)
		inplaceSortRandomNumbers(result)
		fmt.Println("Отсортированные данные:")
		showRandomNumbers(result)
	}
}
