package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	for i:= len(digits)-1; i>=0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func mainFunc() {
	arr := []int{1,2, 3, 9, 9}
	fmt.Println(plusOne(arr))
}