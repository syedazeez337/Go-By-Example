package main

import (
	"fmt"
)

func main() {
	var weightInKgs float64 = 63.45
	var age int = 27

	weightOnMars := weightInKgs * (38.0 / 100.0)
	fmt.Printf("Weight: Earth - %v, Mars - %v\n", weightInKgs, weightOnMars)

	ageOnMars := age * 365 / 687
	fmt.Printf("Age: Earth - %v, Mars - %v\n", age, ageOnMars)
}