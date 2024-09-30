package main

import (
	"fmt"
)

// create a type for a rectangle
type rect struct {
	width int
	length int
}

func (r rect) area() int {
	return r.width * r.length
}

func (r rect) perimeter() int {
	return 2 * (r.length + r.width)
}

func methods() {
	rectangle := rect{width: 34, length: 45}
	fmt.Println("Area is ",rectangle.area())
	fmt.Println("Perimeter is ", rectangle.perimeter())
}