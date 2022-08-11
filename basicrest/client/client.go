package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Schema struct {
	Title   string
	Image   string
	AppName string
}

var URL string = "http://localhost:8003"

func hit() {

	response, prob := http.Get(URL)
	if prob != nil {
		fmt.Printf("Could not reach local server -- %v", prob)
	}
	fmt.Println("response unformatted ", response)
}

func hitlog() {
	newSchema, prob := json.Marshal(Schema{Title: "Banana", Image: "Banana343:v3.0.0", AppName: "Fruits"})
	if prob != nil {
		fmt.Printf("Could not reach marshal schema -- %v", prob)
	}
	URL = "http://localhost:8003/log"
	newReader := bytes.NewReader(newSchema)
	contents, prob := http.Post(URL, "application/json", newReader)
	if prob != nil {
		fmt.Printf("Could not reach local server -- %v %s", prob, URL)
	}

	fmt.Printf("Sending schema to %s\n", URL)
	contentsrep, prob := ioutil.ReadAll(contents.Body)
	if prob != nil {
		fmt.Printf("Could not read response -- %v %s", prob, URL)
	}
	fmt.Printf("%s", contentsrep)
}

func main() {
	hitlog()
}
