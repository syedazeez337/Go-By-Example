package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i * 20
		}
	}()

	go func() {
		for i := 1; i <= 20; i++ {
			fmt.Printf("Value %v\n", <-ch)
		}
	}()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Println("Programme exited")
}
