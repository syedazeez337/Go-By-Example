package main

import (
	"fmt"
)

type person struct {
	name string
	age int
	gender byte
}

func manDetails(name string, age int, gender byte) *person {
	p := person{
		name: name,
		age: age,
		gender: gender,
	}
	return &p
}

func structs() {
	fmt.Println(manDetails("Alex", 12, 'M'))

	womanDetails := person{
		name: "Lua",
		age: 30,
		gender: 'F',
	}

	fmt.Println(womanDetails)
}