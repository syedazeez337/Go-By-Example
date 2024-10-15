package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)
	count := 0

	go func() {
		ch <- "Go message"
	}()

	go func() {
		ch <- "Second go Message"
	}()

	go func() {
		for i := 0; i < 2; i++ {
			<-ch
			count++
		}
		done <- true
	}()

	// Wait for all messages to be processed
	<-done

	fmt.Printf("Total number of strings in the channel: %d\n", count)
}
