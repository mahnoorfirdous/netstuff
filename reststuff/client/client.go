package main

import (
	"fmt"
	"log"
	"net/http"
)

func hit() {
	URL := "http://localhost:8003"
	response, prob := http.Get(URL)
	if prob != nil {
		log.Printf("Could not reach local server -- %v", prob)
	}
	fmt.Println("response unformatted ", response)
}

func main() {
	hit()
}
