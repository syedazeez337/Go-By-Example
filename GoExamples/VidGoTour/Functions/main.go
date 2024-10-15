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

// Variadic function
func sumofNums(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	hello()
	a, b := 4, 2
	fmt.Printf("Add(%d,%d) -> %d\n", a, b, add(a,b))

	mul, div := mulDiv(a, b)
	fmt.Printf("Mul(%d,%d) -> %d, Div(%d, %d) -> %d\n", a, b, mul, a, b, div)

	twoSum := sumofNums(1,2)
	fmt.Printf("Sum of two: %d\n", twoSum)
	threeSum := sumofNums(1,2,3,4,5)
	fmt.Printf("Sum of three: %d\n", threeSum)
	arr := []int {1,2,3,4,5}
	sliceSum := sumofNums(arr...)
	fmt.Printf("Slice sum: %d\n", sliceSum)
}