package main

import (
	"fmt"
)

// anonymous functions
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func anaonfunc() {
	newInt := intSeq()

	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())
}
