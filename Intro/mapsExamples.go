package main

import (
	"fmt"
)

func mapExamples() {
	// making a map
	// Key-value
	mps := map[int] string {1 : "One", 2 : "two"}
	fmt.Println(mps)

	// a new empty map
	maps1 := make(map[string]int)
	fmt.Println(maps1)

	maps1["three"] = 3
	maps1["four"] = 4
	fmt.Println(maps1)
}