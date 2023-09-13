package main

import (
	"fmt"
	"math"
	. "os"
)

func fact(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var a, b, res float64
	var op string

	fmt.Println("Введите операцию через пробел. (пример 1+2 : 1 + 2)")
	fmt.Println("Возведение в степень   (пример два в кубе :  2 ^ 3)")
	fmt.Println("Факториал числа (пример: 1 f 10, 10 исходное число)")
	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		panic(err)
	}
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b != 0 {
			res = a / b
		} else {
			fmt.Println("Hа ноль делить нельзя.")
			return
		}
	case "^":
		res = math.Pow(a, b)
	case "f":
		res = float64(fact(uint(b)))

	default:
		fmt.Println("Операция введена неверно")
		Exit(1)
	}
	if int(res*100)%100 == 0 {
		fmt.Printf("Результат расчета: %v\n", int(res))
	} else {
		fmt.Printf("Результат расчета: %.2f\n", res)
	}
}