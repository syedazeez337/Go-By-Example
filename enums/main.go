package main

import "fmt"

var (
	axe = "axe"
	sword = "sword"
	woodenstick = "woodenstick"
	knife = "knife"
)

func getDamage(weaponType string) (string, int) {
	switch weaponType {
	case axe:
		return axe, 100
	case sword:
		return sword, 90
	case woodenstick:
		return woodenstick, 1
	case knife:
		return knife, 40
	default:
		panic("weapon does not exit")
	}
}

func main() {
	weapon, damage := getDamage(knife)
	fmt.Printf("Damage of a %v is %v\n", weapon, damage)
}