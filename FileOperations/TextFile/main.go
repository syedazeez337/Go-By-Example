package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "file.txt"

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%c ", data)
	fmt.Println()
	fmt.Println(string(data))
}
