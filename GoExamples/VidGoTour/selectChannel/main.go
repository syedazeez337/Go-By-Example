package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	chInt := make(chan int)

	go func() {
		ch1 <- "First message"
	}()

	go func() {
		ch2 <- "Second message"
	}()

	for i:=0; i<10; i++ {
		go func(n int) {
			chInt <- i
		}(i)
	}

	x := "Variable message"
	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	case v := <-chInt:
		fmt.Println(v)
	case ch3 <- x:
		fmt.Println("Wrote", x)
	case <-ch4:
		fmt.Println("Received channel 4, but ignored it")
	}
}