package main

import (
	"fmt"
)

func main() {
	arr := []int {2,3,4,5,6}
	
	ch := make(chan int, len(arr))

	for _, val := range arr {
		go func() {
			ch <- val * 2
		}()
	}

	for i:=0; i < len(arr); i++ {
		fmt.Printf("Value %v\n", <-ch)
	}
}