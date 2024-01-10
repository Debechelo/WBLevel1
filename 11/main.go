package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств

// Intersect находит пересечение двух неупорядоченных множеств
func Intersect(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Итерируем по ключам первого множества
	for key := range set1 {
		// Проверяем, существует ли такой ключ во втором множестве
		if _, exists := set2[key]; exists {
			// Если ключ существует в обоих множествах, добавляем его в результат
			result[key] = true
		}
	}

	return result
}

func main() {
	// Пример двух неупорядоченных множеств
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Находим пересечение
	intersection := Intersect(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", intersection)
}
