package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentence = "This is a sentence"

	fmt.Printf("%v, %v\n", sentence, strings.Contains(sentence, "sentence"))
}