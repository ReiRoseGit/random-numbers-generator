package main

import (
	"flag"
	"fmt"
	"os"

	"random-numbers-generator/generation"
)

func main() {
	// Флаги:
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
		g := new(generation.Generator)
		g.Generate(bound, flows)
		fmt.Println("Неотсортированные данные:")
		g.ShowUnsortedNumbers()
		fmt.Println("Отсортированные данные:")
		g.ShowSortedNumbers()
	}
}
