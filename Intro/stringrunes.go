package main

import (
	"fmt"
	"unicode/utf8"
)

func stringsRunes() {
	const s = "Hello*"

	for i, val := range s {
		fmt.Println(i, val)
	}
	fmt.Println("Length of a string ", len(s))
	fmt.Println("Rune count of a string ", utf8.RuneCountInString(s))
}