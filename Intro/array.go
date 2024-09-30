package main

import (
	"fmt"
)

func array() {
	// declaring an array
	// var arr [5]int
	arr := [5]int{2,3,4,5,6}

	fmt.Println(arr)

	// printing each value using a for loop
	for i, val := range arr {
		fmt.Println(i, val)
	}

	// declaring a 2D array
	twoDarry := [2][3]int {
		{0,1,2},
		{3,4,5},
	}
	fmt.Println(twoDarry)

	// building a 2D array
	var muls[3][3]int
	for i:=0; i<3; i++ {
		for j:=1; j<3; j++ {
			muls[i][j] = i * (j * 3)
		}
	}
	fmt.Println(muls)
}