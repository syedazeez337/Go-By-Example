package main

import (
	"fmt"
)

func forloop() {
	// for as a while loop
	i := 1
	for i <= 10 {
		fmt.Println("Number : ", i)
		i = i + 1
	}

	// for loop
	for i:= 1; i <= 10; i++ {
		fmt.Printf("i = %v\n", i)
	}

	for i := range 20 {
		fmt.Println(i)
	}
}