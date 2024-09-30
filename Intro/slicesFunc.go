package main

import (
	"fmt"
	"slices"
)

func slicesFunc() {
	var sli []int
	fmt.Println("Length of a slice is ", len(sli))

	sli = make([]int, 3)
	fmt.Printf("Length of a slice is %v and capacity is %v\n", len(sli), cap(sli))

	sli[0] = 2
	sli[1] = 3
	sli[2] = 4
	fmt.Println(sli)

	c := make([]string, 3)
	c[0] = "zero"
	c = append(c, "Good")
	c = append(c, "Man", "morning")
	fmt.Println(c)
	fmt.Printf("Length is %v, capacity is %v\n", len(c), cap(c))

	t := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}
	fmt.Println(t)
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	
	s := make([]int, 3)
	fmt.Println(s)

	vals := []string{"g", "h", "i"}
	newT := make([]string, len(vals), (cap(vals)+1) * 2)
	fmt.Println(newT)
	fmt.Printf("length is %v, capacity is %v\n", len(newT), cap(newT))
	newT = vals
	fmt.Println(newT)
	fmt.Printf("length is %v, capacity is %v\n", len(newT), cap(newT))
}