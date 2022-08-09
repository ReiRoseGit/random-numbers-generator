package main

import (
	"flag"
	"os"

	"random-numbers-generator/generation"
)

func main() {
	// Флаги:
	// bound - верхняя граница
	// flows - количество потоков
	var bound, flows int
	flag.IntVar(&bound, "bound", 40, "Верхняя граница (положительное целое число)")
	flag.IntVar(&flows, "flows", 3, "Количество потоков (положительное целое число)")
	flag.Parse()
	if bound < 1 || flows < 1 {
		flag.Usage()
		os.Exit(1)
	} else {
		g := generation.NewGenerator(bound, flows)
		g.Generate()
	}
}
