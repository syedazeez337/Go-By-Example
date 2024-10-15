package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)

	go message("Hello", ch)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %v\n", <-ch)
	}
	fmt.Println("Programme exited")
}

func message(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
