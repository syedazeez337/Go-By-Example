package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Go message"
	}()

	val, ok := <- ch

	fmt.Printf("Value: %s, ok: %v\n", val, ok)
}