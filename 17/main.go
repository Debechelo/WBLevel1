package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 7, 10, 15, 18, 21, 23, 26, 28, 30}

	target := 30

	ind1 := binSearch(arr, target)
	if ind1 < len(arr) && arr[ind1] == target {
		fmt.Printf("Значение %d найдено в массиве на позиции %d.\n", target, ind1)
	} else {
		fmt.Printf("Значение %d не найдено в массиве.\n", target)
	}
}

// Бинарный Поиск
func binSearch(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
	return index
}
