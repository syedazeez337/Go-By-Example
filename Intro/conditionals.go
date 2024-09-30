package main

import (
	"fmt"
)

func conditionals() {
	// conditionals
	num := 2
	if num == 2 {
		fmt.Println("Number is two")
	}

	// if-else conditionals
	num2 := 3
	if num2 != 3 {
		fmt.Println("Number is not equal to 3")
	} else {
		fmt.Println("Number is equal to 3")
	}

	// if-else if-else conditional
	num3 := 4
	if num3 == 2 {
		fmt.Println("Number is two")
	} else if num3 == 4 {
		fmt.Println("Number is 4")
	} else {
		fmt.Println("Number can be anything..")
	}
}