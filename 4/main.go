package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из
// канала и выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

const (
	durationSeconds = 2 // Время обработки worker
	numWorkers      = 3 // Задаем количество воркеров
)

// Функция writer записывает данные в канал.
func writer(ch chan<- int, wg *sync.WaitGroup, signalCh <-chan os.Signal) {
	defer wg.Done()
	for i := 1; ; i++ {
		select {
		case <-signalCh:
			// Закрываем канал ch чтобы worker-ы не продолжали работать с ним
			close(ch)
			fmt.Printf("Writer finished.\n")
			return
		case ch <- i:
		}
	}
}

// Функция worker читает данные из канала и выводит их в stdout.
func worker(workerID int, ch <-chan int, wg *sync.WaitGroup, signalCh <-chan os.Signal) {
	defer wg.Done()
	for {
		select {
		case <-signalCh:
			fmt.Printf("Worker %d finished.\n", workerID)
			// Завершаем функцию
			return
		case data, ok := <-ch:
			if ok {
				fmt.Printf("Worker %d: Received data %d\n", workerID, data)
				time.Sleep(time.Duration(durationSeconds) * time.Second)
			}
		}
	}
}

func main() {
	ch := make(chan int, numWorkers)    // Создаем канал с буфером для записи данных.
	signalCh := make(chan os.Signal, 1) // Канал для получения сигналов от системы.

	var wg sync.WaitGroup

	// Запускаем горутину для записи данных в канал.
	wg.Add(1)
	go writer(ch, &wg, signalCh)

	// Запускаем горутины для воркеров.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg, signalCh)
	}

	// Ожидаем получения сигнала завершения от системы.
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
	close(signalCh) // Закрываем канал для сигнала завершения.

	// Ожидаем завершения всех горутин.
	wg.Wait()
}
