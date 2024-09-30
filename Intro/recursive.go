package main

import (
	"fmt"
)

func recursive() {
	fmt.Println(facto(5))
	fmt.Println(facto(6))
	fmt.Println(facto(9))
	fmt.Println(facto(20))
}

// fibonacci function
func facto(n int) int {
	if n == 0 {
		return 1
	}
	return n * facto(n - 1)
}