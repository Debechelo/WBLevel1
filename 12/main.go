package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func main() {
	// Последовательность строк
	stringsSequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	stringSet := make(map[string]struct{})

	// Добавляем уникальные строки в множество
	for _, str := range stringsSequence {
		stringSet[str] = struct{}{}
	}

	// Выводим строки
	for str := range stringSet {
		fmt.Println(str)
	}
}
