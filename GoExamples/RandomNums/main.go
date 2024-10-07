package main

import (
	"fmt"
	"math/rand"
)


func main() {
	var num = rand.Intn(10) + 1
	// to get a random integer from 1 - 10
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}