package main

import (
    "fmt"
    "strings"
)

func main() {
    var input string
    fmt.Print("Введите арифметическое выражение: ")
    fmt.Scanln(&input)

    // Разделяем ввод на операнды и операцию
    values := strings.Split(input, " ")
    if len(values) != 3 {
        fmt.Println("Ошибка: неверный формат ввода")
        return
    }

    // Проверяем, что операнды являются допустимыми числами
    a, valid := convertToArabic(values[0])
    if !valid {
        fmt.Println("Ошибка: неверный формат числа")
        return
    }

    b, valid := convertToArabic(values[2])
    if !valid {
        fmt.Println("Ошибка: неверный формат числа")
        return
    }

    // Проверяем, что оба операнда являются либо арабскими, либо римскими числами
    isArabic := isArabicNumber(values[0]) && isArabicNumber(values[2])
    isRoman := isRomanNumber(values[0]) && isRomanNumber(values[2])

    if !(isArabic || isRoman) {
        fmt.Println("Ошибка: числа должны быть одного типа (арабские или римские)")
        return
    }

    // Вычисляем и выводим результат
    var result int
    switch values[1] {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        result = a / b
    default:
        fmt.Println("Ошибка: недопустимая операция")
        return
    }

    if isRoman {
        // Преобразуем результат в римское число
        romanResult, err := convertToRoman(result)
        if err != nil {
            fmt.Println("Ошибка: неверный результат для римских чисел")
            return
        }
        fmt.Println("Результат:", romanResult)
    } else {
        fmt.Println("Результат:", result)
    }
}

func convertToArabic(roman string) (int, bool) {
    // Проверяем, что число является римским
    if !isRomanNumber(roman) {
        return 0, false
    }

    // Определяем значение каждого символа римского числа
    romanMap := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    result := 0
    prev := 0

    // Проходим по всем символам числа и вычисляем его значение
    for _, digit := range roman {
        value := romanMap[digit]
        if value > prev {
            result += value - 2*prev
        } else {
            result += value
        }
        prev = value
    }

    return result, true
}

func convertToRoman(arabic int) (string, error) {
    // Проверяем, что число находится в допустимом диапазоне
    if arabic <= 0 {
        return "", fmt.Errorf("неверный результат для римских чисел")
    }

    // Определяем значения для каждого разряда римских чисел
    romanMap := map[int]string{
        1000: "M",
        900:  "CM",
        500:  "D",
        400:  "CD",
        100:  "C",
        90:   "XC",
        50:   "L",
        40:   "XL",
        10:   "X",
        9:    "IX",
        5:    "V",
        4:    "IV",
        1:    "I",
    }

    result := ""
    // Проходим по всем разрядам и добавляем соответствующие символы римского числа
    for arabic > 0 {
        for value, symbol := range romanMap {
            if arabic >= value {
                result += symbol
                arabic -= value
                break
            }
        }
    }

    return result, nil
}

func isArabicNumber(number string) bool {
    // Проверяем, что число является арабским
    _, err := convertToArabic(number)
    return err == nil
}

func isRomanNumber(number string) bool {
    // Проверяем, что число является римским
    romanSymbols := []string{"I", "V", "X", "L", "C", "D", "M"}

    for _, symbol := range number {
        found := false
        for _, roman := range romanSymbols {
            if string(symbol) == roman {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }

    return true
}
