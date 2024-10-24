package main

import (
	"fmt"
)

func main() {

	var a [4]int

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4

	fmt.Println("Array>", a)

	b := [5]int{21, 22, 32, 45, 67}
	fmt.Println("Array 2>", b)

	// slice declaration
	c := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("Slice>", c)
	fmt.Println("Indexed slice>", c[2:5])

	d := make([]int, 5)
	fmt.Println("Empty slice>", d)
	d = append(d, 2)
	d = append(d, 3)
	fmt.Println("Refined slice>", d)

	phoneBook := map[string]int{
		"A": 123,
		"C": 456,
		"B": 789,
	}

	fmt.Println(phoneBook)
}
