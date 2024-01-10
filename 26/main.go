package main

import "fmt"

// Функция, которая проверяет, что все символы в строке уникальные
func areAllCharactersUnique(input string) bool {
	// Используем map для отслеживания уникальных символов
	characterSet := make(map[rune]bool)

	// Перебираем символы в строке
	for _, char := range input {
		// Если символ уже присутствует в наборе, возвращаем false
		if characterSet[char] {
			return false
		}
		// Добавляем символ в набор
		characterSet[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	testString1 := "abcd"      // true
	testString2 := "abCdefAaf" // false
	testString3 := "aabcd"     // false

	fmt.Printf("%s: %t\n", testString1, areAllCharactersUnique(testString1))
	fmt.Printf("%s: %t\n", testString2, areAllCharactersUnique(testString2))
	fmt.Printf("%s: %t\n", testString3, areAllCharactersUnique(testString3))
}
