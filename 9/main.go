package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
// из массива, во второй — результат операции x*2, после чего данные из
// второго канала должны выводиться в stdout.

func main() {
	// Создаем каналы
	inputChan := make(chan int)
	outputChan := make(chan int)

	var wg sync.WaitGroup

	// Горутина для записи чисел в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(inputChan) // Закрываем канал после завершения записи

		numbers := []int{1, 2, 3, 4, 5}

		for _, num := range numbers {
			inputChan <- num
		}
	}()

	// Горутина для умножения чисел из первого канала и записи во второй канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(outputChan) // Закрываем канал после завершения записи

		for num := range inputChan {
			result := num * num
			outputChan <- result
		}
	}()

	// Горутина для вывода результатов из второго канала
	wg.Add(1)
	go func() {
		defer wg.Done()

		for result := range outputChan {
			fmt.Println(result)
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
}
