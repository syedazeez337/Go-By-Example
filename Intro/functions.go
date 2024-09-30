package main

import (
	"fmt"
)

func functions() {
	fmt.Println("Add(2,3) = ", addTwo(2,3))
	fmt.Println("Add(2,3,4) = ", addThree(2,3,4))
	a, b := addAndMul(3,4)
	fmt.Printf("Add(3,4) = %v, Mul(3,4) = %v\n", a, b)
}

// an addTwo function
func addTwo(x int, y int) int {
	return x + y
}

// an addthree function
func addThree(x int, y int, z int) int {
	return x + y + z
}

// multiple return values
func addAndMul(x int, y int) (int, int) {
	add := x + y
	mul := x * y
	return add, mul
}