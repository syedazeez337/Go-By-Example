package main

import (
	"fmt"
)

type World struct{

}

func (w *World) String() string {
	return "World"
}

func main() {
	fmt.Println("Hello ", new(World))
}