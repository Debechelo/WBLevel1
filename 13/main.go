package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 1, 5

	// 1
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

	// 2
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

	// 3
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Обмен значений
	a, b = b, a

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

}
