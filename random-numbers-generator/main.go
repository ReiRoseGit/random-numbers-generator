package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

// получает случайное число, не превосходящее верхнюю границу bound
func getRandomNumber(bound int, channel chan int) {
	channel <- 1 + rand.Intn(bound)
}

// функция-диспетчер
// запускает функцию получения случайного числа и его фильтрации до тех пор, пока не сгенерируется нужное кол-во элементов
// для каждого потока получения и фильтрации случайного значения создается свой канал: channels[i]
// по этому каналу передается сначала случайно сгенерированное случайное число (вызов getRandomNumber)
// затем в анонимной функции происходит фильтрация этого значения:
// - извлечение значения из канала
// - фиксирование значения словаря с помощью мьютекса
// - сравнение полученного по каналу значения со значениями словаря usedNumbers
// - добавления значения в словарь использованных чисел (usedNumbers) и в срез-результат (result) в случае уникальности числа
func getAllRandomNumbers(bound, flows int) []int {
	rand.Seed(time.Now().UnixNano())
	var mutex sync.Mutex
	usedNumbers := make(map[int]bool)
	result := []int{}
	channels := []chan int{}
	for i := 0; i < flows; i++ {
		channels = append(channels, make(chan int))
	}
	for {
		if len(usedNumbers) == bound {
			break
		}
		for i := 0; i < flows; i++ {
			go getRandomNumber(bound, channels[i])
			// функция-фильтратор
			// выбирает допустимое значение, опираясь на словарь использованных чисел
			// новое число не должно принадлежать словарю usedNumbers
			go func(mutex *sync.Mutex, channel chan int) {
				number := <-channel
				mutex.Lock()
				if _, ok := usedNumbers[number]; !ok {
					usedNumbers[number] = true
					result = append(result, number)
				}
				mutex.Unlock()
			}(&mutex, channels[i])
		}
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
	flag.IntVar(&bound, "bound", 10, "Верхняя граница (положительное целое число)")
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
