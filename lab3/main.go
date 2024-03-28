package main

import (
	"calculator/calculations"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var logg bool
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Pass an argument to the function")
		os.Exit(1)
	}
	n, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	flag.BoolVar(&logg, "log", true, "enable logging")
	flag.Parse()

	calculations.Calculate(n, logg)

	// тут вывод с помощью пакеты из репозитория
}
