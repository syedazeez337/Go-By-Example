package main

func pow(x int) int {
	r := x
	for r * r > x {
		r = (r + x / r) / 2
	}
	return r
}