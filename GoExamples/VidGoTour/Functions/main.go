package main

import (
	"fmt"
)

// normal functions call with no return type
func hello() {
	fmt.Println("Hello function")
}

// function with return type
func add(a int, b int) int {
	return a + b
}

// function with multiple return type values
func mulDiv(a int, b int) (int, int) {
	return a * b, a / b
}

func main() {
	hello()
	a, b := 4, 2
	fmt.Printf("Add(%d,%d) -> %d\n", a, b, add(a,b))

	mul, div := mulDiv(a, b)
	fmt.Printf("Mul(%d,%d) -> %d, Div(%d, %d) -> %d\n", a, b, mul, a, b, div)
}