package main

import (
	"fmt"
)

func constants() {
	// declaring a const variable
	const value = 493284

	// const values cannot be changed in the programme
	const s string = "a constant"

	fmt.Printf("constant values %v, %v\n", value, s)
}