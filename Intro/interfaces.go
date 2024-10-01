package main

import (
	"fmt"
	"math"
)

func square(x float64) float64 {
	return x * x
}

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	length float64
	breadth float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perim() float64 {
	return 2 * (r.length + r.breadth)
}

func (c circle) area() float64 {
	return math.Pi * square(c.radius)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func interfaces() {
	rec := rectangle{length: 4, breadth: 3}
	cir := circle{radius: 5}

	measure(rec)
	measure(cir)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}