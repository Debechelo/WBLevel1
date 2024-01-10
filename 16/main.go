package main

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	// Создаем неотсортированный срез
	arr := []int{3, 6, 4, 1, 35, 9, 2, 6, 8, 1, 7}

	// Используем встроенную сортировку
	sort.Ints(arr)

	fmt.Println(arr)

	// или сами задаем функцию по которой сортируем числа
	// сортировка в по убыванию
	arr = []int{3, 6, 4, 1, 35, 9, 2, 6, 8, 1, 7}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	fmt.Println(arr)
}
