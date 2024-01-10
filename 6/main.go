package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// 1. Использование контекста
// Горутина работает с контекстом и завершает
// выполнение при сигнале отмены контекста.
func methodContext() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) { // Горутина
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // Если контекст становился, останавливаем горутину
				fmt.Println("Goroutine: Received stop signal. Exiting.")
				return
			default:
				// Выполнение работы горутины
				fmt.Println("Goroutine: Working...")
				time.Sleep(time.Second)
			}
		}
	}(ctx, &wg)

	time.Sleep(3 * time.Second)
	cancel() // Отправляем сигнал об отмене контекста

	// Ожидаем завершения всех горутин
	wg.Wait()
}

// 2. Использование канала для сигнала завершения:
// Горутина опрашивает канал на предмет сигнала о завершении
func methodChanel() {
	stopCh := make(chan struct{}) // канал сигнализации завершения
	var wg sync.WaitGroup

	wg.Add(1)
	go func(ch <-chan struct{}, wg *sync.WaitGroup) { // Горутина
		defer wg.Done()

		for {
			select {
			case <-ch: // Как только что то попадет в канал или канал закрыт, завершаем работу
				fmt.Println("Goroutine: Received stop signal. Exiting.")
				return
			default:
				// Выполнение работы горутины
				fmt.Println("Goroutine: Working...")
				time.Sleep(time.Second)
			}
		}
	}(stopCh, &wg)

	// После какого-то времени отправляем сигнал о завершении
	time.Sleep(3 * time.Second)
	stopCh <- struct{}{} // Запись
	close(stopCh)        // или Закрытие канала

	// Ожидаем завершения всех горутин
	wg.Wait()
}

// 3. Использование runtime.Goexit():
// Но это жесткий способ так как runtime.Goexit()
// не пощадит горутину и закроет ее немедленно
// без выполнения defer или очистки каких либо ресурсов
func methodExit() {
	defer fmt.Println("Goroutine: Stop!")
	go func() {
		for i := 0; ; i++ {
			// Выполнение работы горутины
			fmt.Println("Goroutine: Working...")
			time.Sleep(time.Second)
			if i == 3 { // В определеный момент
				runtime.Goexit()
			}
		}
	}()

	// Подождем некоторое время
	time.Sleep(5 * time.Second)
}

func main() {
	methodContext()
	methodChanel()
	methodExit()
}
