package main

import (
	"fmt"
)

func main() {
	const SPEEDLIGHT = 299792 // km/s

	var distance = 56000000
	fmt.Println(distance / SPEEDLIGHT, "Seconds")

	distance = 401000000
	fmt.Println(distance / SPEEDLIGHT, "Seconds")
}