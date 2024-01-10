package main

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

// SetBit устанавливает i-й бит в num в значение bit
func SetBit(num int64, i uint, bit int) int64 {
	mask := int64(1) << i // Создаем маску, устанавливая i-й бит в 1
	if bit == 1 {
		return num | mask // Устанавливаем i-й бит в 1, используя побитовое ИЛИ
	}
	return num &^ mask // Устанавливаем i-й бит в 0, используя побитовое AND NOT
}

func main() {
	var num int64
	var i uint
	var bit int

	// Вводим значения
	fmt.Print("Введите число (int64): ")
	fmt.Scan(&num)

	fmt.Print("Введите номер бита(начиная с 0): ")
	fmt.Scan(&i)

	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&bit)
	if bit != 0 && bit != 1 {
		fmt.Print("Ошибка!!")
		return
	}

	// Устанавливаем i-й бит в 1 или 0
	result := SetBit(num, i, bit)

	// Выводим результат
	fmt.Printf("Результат: %d\n", result)
}
