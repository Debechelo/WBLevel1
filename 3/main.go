package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

// Функция calculateSquare рассчитывает квадрат числа и отправляет результат в канал.
func calculateSquare(number int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	result := number * number
	ch <- result
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup              // Создаем группу для ожидания завершения всех горутин.
	ch := make(chan int, len(numbers)) // Создаем канал для передачи результатов.

	for _, num := range numbers {
		wg.Add(1)

		// Запускаем горутину для каждого числа.
		go calculateSquare(num, &wg, ch)
	}

	// Горутина для закрытия канала после завершения всех горутин.
	go func() {
		wg.Wait() // Ждем, пока все горутины завершат свою работу.
		close(ch)
	}()

	// Читаем из канала и суммируем результаты.
	var sum int
	for result := range ch {
		sum += result
	}

	fmt.Println("The sum of the squares of the numbers:", sum)
}
