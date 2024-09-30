package main

import (
	"fmt"
)

func zerovalue(ivalue int) {
	ivalue = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func pointers() {
	i := 1
	fmt.Println(i)

	zerovalue(i)
	fmt.Println(i)

	zeroPtr(&i)
	fmt.Println(i)

	fmt.Println(&i)
}