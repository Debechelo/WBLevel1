package main

import (
	"fmt"
	"sync"
)

// Counter Структура
type Counter struct {
	count int
	mu    sync.Mutex
}

// Метод для инкрементации счетчика
func (c *Counter) increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	// Но тут можно изпользовать
	// atomic.AddInt64(&c.count, 1)
	// Без мьютекса
}

func worker(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter.increment()
	}
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&counter, &wg)
	}

	wg.Wait()

	fmt.Println("Итоговое значение счетчика:", counter.count)
}
