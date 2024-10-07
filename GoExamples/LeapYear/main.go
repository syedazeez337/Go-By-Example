package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Printf("Enter a year: ")
	fmt.Scanf("%d", &year)

	// check if it is leap year or not
	var leap = year%400 == 0 || (year%4==0 && year%100!=0)

	if leap {
		fmt.Printf("Year %v is a Leap year\n", year)
	} else {
		fmt.Println("Not a leap year")
	}
}