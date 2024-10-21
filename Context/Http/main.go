package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	startTime := time.Now()
	response, err := http.Get("https://www.google.com/")

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	endTime := time.Since(startTime)

	fmt.Println("Response:", response.Status)

	fmt.Println("Time taken: ", endTime)
}
