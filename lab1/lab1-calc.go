package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func check_num(num string) bool {
	if !(isNumeric(num)) {
		fmt.Println("Некоректнный ввод. Пожалуйства, введите числовое значение.")
		return true
	} else {
		return false
	}
}

func check_sym(sym string) bool {
	if (len(sym) == 1) && ((strings.Contains(sym, "+")) || (strings.Contains(sym, "-")) || (strings.Contains(sym, "*")) || (strings.Contains(sym, "/"))) {
		return false
	} else {
		fmt.Println("Некоректнный ввод. Пожалуйства, выберите операцию.")
		return true
	}
}

func main() {
	var (
		first     string
		procedure string
		second    string
	)
	for {
		fmt.Println("Введите первое число: ")
		fmt.Scanln(&first)
		for {
			if check_num(first) {
				fmt.Scanln(&first)
			} else {
				break
			}
		}
		fmt.Println("Выберите операцию(+,-,*,/): ")
		fmt.Scanln(&procedure)
		for {
			if check_sym(procedure) {
				fmt.Scanln(&procedure)
			} else {
				break
			}
		}
		fmt.Println("Введите второе число: ")
		fmt.Scanln(&second)
		for {
			if check_num(second) {
				fmt.Scanln(&second)
			} else {
				break
			}

		}

		if procedure == "+" {
			first, err := strconv.Atoi(first)
			second, err := strconv.Atoi(second)
			sum := int(first) + int(second)
			fmt.Println("Результат: ", sum)
			if err != nil {
				log.Fatal(err)
			}
		}

		if procedure == "-" {
			first, err := strconv.Atoi(first)
			second, err := strconv.Atoi(second)
			div := int(first) - int(second)
			fmt.Println("Результат: ", div)
			if err != nil {
				log.Fatal(err)
			}
		}
		if procedure == "*" {
			first, err := strconv.Atoi(first)
			second, err := strconv.Atoi(second)
			sum := int(first) * int(second)
			fmt.Println("Результат: ", sum)
			if err != nil {
				log.Fatal(err)
			}
		}
		if procedure == "/" {
			first, err := strconv.Atoi(first)
			second, err := strconv.Atoi(second)
			if second == 0 {
				fmt.Println("Деление на ноль невозможно")
				continue
			}
			sum := int(first) / int(second)
			fmt.Println("Результат: ", sum)
			if err != nil {
				log.Fatal(err)
			}

		}

	}
}
