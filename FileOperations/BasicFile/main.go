package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	fileName := flag.String("go", "file.go", "A go file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("File cannot be found")
	} else {
		fmt.Println("Success", file)
	}
}