package main

import (
	"fmt"
	"strings"
)

// Функция, которая переворачивает слова в строке
func reverseWords(input string) string {
	words := strings.Fields(input)
	length := len(words)

	// Инвертирование порядка слов
	for i := 0; i < length/2; i++ {
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	inputString := "snow dog sun"
	reversedWordsString := reverseWords(inputString)

	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Строка с инвертированными словами: %s\n", reversedWordsString)
}
