package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	startTime := time.Now()
	wg.Add(1)
	go httpResponse("https://www.google.com/")

	wg.Wait()

	wg.Add(1)
	go httpResponse("https://www.google.com/")

	wg.Wait()

	wg.Add(1)
	go httpResponse("https://www.google.com/")

	wg.Wait()
	endTime := time.Since(startTime)
	fmt.Println("Time taken: ", endTime)

}

func httpResponse(web string) *http.Response {
	defer wg.Done()
	resp, err := http.Get(web)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Response: ", resp.Status)
	return resp
}
