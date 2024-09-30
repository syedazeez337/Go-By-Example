package main

import (
	"fmt"
)

func variadicFuncs() {
	buildAsentence("what ", "is ", "today?")
}

// variadic function 
// build a sentence
func buildAsentence(name... string) {
	total := ""
	for _, words := range name {
		total += words
	}
	fmt.Println(total)
}