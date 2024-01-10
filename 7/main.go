package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

// 1. Map использование Mutex
func WriteMap(mp map[string]string, key, value string, mu *sync.Mutex) {
	mu.Lock() // Блокируем мьютекс перед доступом к map
	mp[key] = value
	mu.Unlock() // Освобождаем мьютекс
}

func MapWithMutex() {
	const count = 5
	mp := make(map[string]string, count)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(mp map[string]string, index int, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			key := fmt.Sprintf("Key%d", index)
			value := fmt.Sprintf("Value%d", index)

			// Конкурентная запись в map
			WriteMap(mp, key, value, mu)
		}(mp, i, &wg, &mu)
	}

	// Подождем завершения горутин
	wg.Wait()

	// Выведем результаты
	fmt.Println("Final Map Contents:")
	for key, value := range mp {
		fmt.Printf("%s -> %s\n", key, value)
	}
}

// 2. sync.Map
func WriteSyncMap(mp *sync.Map, key, value string) {
	mp.Store(key, value)
}

func SyncMap() {
	const count = 5
	var mp sync.Map
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup, mp *sync.Map) {
			defer wg.Done()
			key := fmt.Sprintf("Key%d", index)
			value := fmt.Sprintf("Value%d", index)

			// Конкурентная запись в sync.Map
			WriteSyncMap(mp, key, value)
		}(i, &wg, &mp)
	}

	// Подождем завершения горутин
	wg.Wait()

	// Выведем результаты
	fmt.Println("Final Map Contents:")
	mp.Range(func(key, value interface{}) bool {
		fmt.Printf("%s -> %s\n", key, value)
		return true
	})
}

func main() {
	MapWithMutex()
	SyncMap()
}
