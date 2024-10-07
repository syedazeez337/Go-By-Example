package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myGuess = 34

	
	for {
		var randNum = rand.Intn(100) + 1
		if randNum > myGuess {
			fmt.Printf("Number %v is too high\n", randNum)
		} else if randNum < myGuess {
			fmt.Printf("Number %v is too low\n", randNum)
		} else {
			fmt.Printf("Number has matched Computer guess : %v, My guess: %v\n", randNum, myGuess)
			break
		}
	}
}