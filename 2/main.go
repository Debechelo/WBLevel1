package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов
//чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// Функция calculateSquare рассчитывает квадрат числа и выводит результат в stdout.
func calculateSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := number * number
	fmt.Printf("%d squared is %d\n", number, result)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // Создаем группу для ожидания завершения всех горутин.

	for _, num := range numbers {
		wg.Add(1) // Увеличиваем счетчик группы для каждой горутины.

		// Запускаем горутину для каждого числа.
		go calculateSquare(num, &wg)
	}

	// Ждем, пока все горутины завершат свою работу.
	wg.Wait()
}
