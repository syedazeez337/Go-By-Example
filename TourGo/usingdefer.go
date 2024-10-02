package main

import (
	"fmt"
)

func usingDefer() {
	defer fmt.Println("world2")

	fmt.Print("Hello ")
	fmt.Print("world ")
}