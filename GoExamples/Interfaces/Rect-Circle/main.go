package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	length, width float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	r := rect{length: 23, width: 12}

	c := circle{radius: 34}

	fmt.Println(r.area())
	fmt.Println(c.area())

}
