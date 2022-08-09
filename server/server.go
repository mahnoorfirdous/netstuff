package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//TODO: Add routing implementations and imports
)

type Notebook struct {
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Width        string `json:"width"`
}

var notebooks []Notebook

func homepage(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "<h1> Welcome! You are at the homepage </h1>\n")
	fmt.Println("Hit Endpoint /: home ")
	fmt.Fprintf(rw, "<p> Sorry this isn't a fashionable page </p>")
}

func handle_requests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/notebooks", getNotebooks)
	http.HandleFunc("/log", logMessages)
	log.Fatal(http.ListenAndServe(":8003", nil)) //TODO: Secure with TLS
}

func getNotebooks(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Hit Endpoint /: notebooks")
	json.NewEncoder(rw).Encode(notebooks)
}

func logMessages(rw http.ResponseWriter, req *http.Request) {
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("server: I don't get it %s\n", err)
	}
	log.Printf("Received %s\n", contents)
}

func main() {
	handle_requests()

}
