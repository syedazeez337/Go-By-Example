package main

import (
	"fmt"
	"time"
)

func main() {

	arr := []int {22,23,34,44,52}

	go func() {
		fmt.Println("First function")
	}()

	go func() {
		fmt.Println("Second function")
	}()

	for i := 0; i < 5; i++ {
		go countGoroutine(i)
	}

	for _, elem := range arr {
		go countGoroutine(elem)
	}

	time.Sleep(time.Second)
	fmt.Println("Programme exited")
}

func countGoroutine(n int) {
	fmt.Println("Goroutine message", n)
}
