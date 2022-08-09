package main

import (
	"fmt"
	server "samplerest/server"
	token "samplerest/token"
)

func main() {
	newconfig := server.Config{TokenSymmetricKey: token.RandomString(32)}
	newserver, err := server.NewServer(newconfig)
	if err != nil {
		fmt.Println("Unable to create: ", err)
	}
	err = newserver.Start(":8090")
	if err != nil {
		fmt.Println("Unable to start: ", err)
	}

}
