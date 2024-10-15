package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("First function")
	}()

	go func() {
		fmt.Println("Second function")
	}()

	for i := 0; i < 10; i++ {
		go countGoroutine(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Programme exited")
}

func countGoroutine(n int) {
	fmt.Println("Goroutine message", n)
}
