package main

import (
	"fmt"
)

func main() {
	const hoursPerDay = 24
	// speed = distance / time

	// speed in kms/hr
	var distance = 56000000 // km
	var time = 28 * hoursPerDay // hr

	var speed = distance / time

	fmt.Printf("Speed: %v km/h\n", speed)
}