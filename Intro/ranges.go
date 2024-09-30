package main

import (
	"fmt"
)

func ranges() {
	arr := []int {1,2,3,4}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	nums := map[string]int {"one":1,"two":2,"three":3}

	for key, value := range nums {
		fmt.Println(key, value)
	}
}