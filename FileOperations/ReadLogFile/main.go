package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "log.txt"

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
