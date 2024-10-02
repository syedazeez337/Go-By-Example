package main

import (
	"fmt"
)

func addArr() {
	arr1 := []int{1,2,3,4,5}
	fmt.Println(addfunc(arr1))
}

func addfunc(arr [] int) int {
	sum := 0
	for _, num := range arr {
		sum = sum + num
	}
	return sum
}