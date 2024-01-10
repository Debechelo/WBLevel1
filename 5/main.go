package main

import (
	"fmt"
	"sync"
	"time"
)

const durationSeconds = 5

// sender отправляет значения в канал каждую секунду в течение duration секунд.
// Завершает работу, когда проходит указанное количество секунд.
func sender(ch chan<- int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Отправляем значения в канал каждую секунду
	for i := 1; ; i++ {
		select {
		case ch <- i:
		case <-done:
			fmt.Println("Sender: Done signal received. Exiting.")
			return
		}
		time.Sleep(time.Second) // Пауза между отправкой значениями
	}
}

// receiver читает значения из канала и выводит их в stdout.
// Завершает работу при закрытии канала.
func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Читаем значения из канала
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed. Receiver finished.")
				return
			}
			fmt.Printf("Received value: %d\n", value)
		}
	}
}

func wait(ch chan<- int, done chan<- struct{}, duration time.Duration) {
	for {
		select {
		case <-time.After(duration):
			// Закрываем каналы, чтобы завершить работу горутин
			close(ch)
			close(done)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Запускаем горутину-отправитель
	go wait(ch, done, time.Duration(durationSeconds)*time.Second)

	wg.Add(1)
	go sender(ch, done, &wg)

	// Запускаем горутину-получатель
	wg.Add(1)
	go receiver(ch, &wg)

	// Ожидаем завершения горутин по истечению времени
	//time.Sleep(time.Duration(durationSeconds) * time.Second)

	// Ожидаем завершения всех горутин
	wg.Wait()
}
