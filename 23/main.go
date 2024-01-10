package main

import "fmt"

// Функция для удаления i-го элемента из слайса
func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	// Удаляем элемент с индексом 2
	indexToRemove := 2
	slice = removeElement(slice, indexToRemove)

	fmt.Println(slice) // [1 2 4 5]
}
