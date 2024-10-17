package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := flag.String("go", "file.go", "A go file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		//fmt.Println("File cannot be found")
		log.Fatal(err)
	} else {
		fmt.Println("Success", file)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %q\n", data, data[:count])
	fmt.Println(count)

	for i:=0; i<len(data); i++ {
		fmt.Printf("%q ", data[i])
	}
	fmt.Println()
}
