package main

import (
	"fmt"
)

func switchstatement() {
	// days of a week
	dayOfaWeek(2)
	dayOfaWeek(4)

}

func dayOfaWeek(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid number")
	}
}