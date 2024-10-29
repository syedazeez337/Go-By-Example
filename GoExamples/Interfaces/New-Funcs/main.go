package main

import "fmt"

type dummy struct {
	word string
	num  int
}

type anothetDummy struct {
	wat int
}

type dummies interface {
	convert()
	display()
}

func (d dummy) convert() {
	fmt.Println(d.num, d.word)
}

func (a anothetDummy) display() {
	fmt.Println(a.wat)
}

func main() {
	d := dummy{num: 43, word: "What"}
	a := anothetDummy{wat: 76}

	d.convert()
	a.display()
}
