package main

import (
	"fmt"
	"math/rand"
	"time"
)

func message(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go message("Hello")
	fmt.Println("Programme waiting")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Println("Programme exited")
}
