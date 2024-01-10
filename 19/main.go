package main

import (
	"fmt"
)

// Функция, которая переворачивает строку с учетом символов Unicode
func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	// Инвертирование символов в массиве рун
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	return string(runes)
}

func main() {
	inputString := "главрыбай"
	reversedString := reverseString(inputString)

	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
