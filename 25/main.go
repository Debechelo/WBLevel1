package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Printf("Sleep\n")
	sleep(5 * time.Second)
	fmt.Printf("Wake up.")
}
