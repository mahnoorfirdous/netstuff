package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/golang/snappy"
)

const decodeReadLimit = 32 * 1024 * 1024

func main() {

	http.HandleFunc("/receive", dumpHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))

}

func dumpHandler(w http.ResponseWriter, r *http.Request) {
	dumpmessage, err := io.ReadAll(io.LimitReader(r.Body, decodeReadLimit))
	if err != nil {
		fmt.Fprintf(w, err.Error()+"Could not read body")
		log.Println("Could not ready body")
	}

	dumped, err := snappy.Decode(nil, dumpmessage)

	if err != nil {
		fmt.Fprintf(w, err.Error()+": Could not decode")
		log.Println("Decode problem with snappy")
	}

	fmt.Fprintf(w, string(dumped))
	log.Println(string(dumped))
}
