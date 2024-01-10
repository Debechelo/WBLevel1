package main

import (
	"fmt"
	"math"
)

// Перемножение
func multiply(a, b float64) float64 {
	return a * b
}

// Деление
func divide(a, b float64) float64 {
	return b / a
}

// Сложение
func add(a, b float64) float64 {
	return a + b
}

// Вычитание
func subtract(a, b float64) float64 {
	return b - a
}

func main() {
	// Задаем две числовые переменные a и b
	a := math.Pow(2, 25)
	b := math.Pow(2, 22) + 1

	multiplicationResult := multiply(a, b)
	fmt.Printf("Умножение: %.0f * %.0f = %.0f\n", a, b, multiplicationResult)

	divisionResult := divide(a, b)
	fmt.Printf("Деление: %.0f / %.0f = %.2f\n", b, a, divisionResult)

	additionResult := add(a, b)
	fmt.Printf("Сложение: %.0f + %.0f = %.0f\n", a, b, additionResult)

	subtractionResult := subtract(a, b)
	fmt.Printf("Вычитание: %.0f - %.0f = %.0f\n", b, a, subtractionResult)
}
