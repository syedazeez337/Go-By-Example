package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// slice of spacelines
	spacelines := []string {
		"Virgin Galactic",
		"SpaceX",
		"Space Adventure",
		"JAXA",
		"ISRO",
	}

	// trips
	trips := []string {
		"Round-trip",
		"One-way",
	}

	fmt.Println("Spaceline         Days Trip-type   Price(in millions)")
	fmt.Println("=====================================================")
	for i:=0; i<10; i++ {
		spaceRand := rand.Intn(len(spacelines))
		tripRand := rand.Intn(len(trips))
		daysRand := rand.Intn(15) + 15
		costRand := rand.Intn(30) + 20

		spaceLineValue := chooseRandom(spacelines, spaceRand)
		daysValue := randDays(daysRand)
		tripValue := chooseRandom(trips, tripRand)
		costValue := costing(costRand, daysValue)
		fmt.Printf("%16v %4d %10v  $%4d\n",spaceLineValue, daysValue, tripValue, costValue)
	}
	
}

// choosing a random spacelines
func chooseRandom(sp []string, count int) string {
	return sp[count]
}

// random days
func randDays(t int) int {
	distance := 62100000 // km
	speed := t
	time := (distance / speed) / (60 * 60 * 24)
	// fmt.Println(time)
	return time
}

// cost per spaceline
func costing(rCost int, days int) int {
	if(days > 30) {
		return rCost
	} else {
		return double(rCost)
	}
}

// double function
func double(n int) int {
	return n * 2
}
