package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
}

func main() {
	start := time.Now()
	fetch("http://www.google.com")
	fmt.Println(time.Since(start))
}
