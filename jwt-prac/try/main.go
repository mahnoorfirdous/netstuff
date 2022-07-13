package main

import (
	"fmt"
	"time"

	. "jwtry"
)

func main() {
	maker, err := NewJWTMake(RandomString(32)) //make a random token using a random alphabet string
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	username := RandomString(4)  //make a random username of 4 characters
	timespan := time.Second * 10 // we want the token to last for a minute

	issuedAttime := time.Now()
	expiredAttime := issuedAttime.Add(timespan) //the deadline
	fmt.Println(expiredAttime)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	token, err := maker.CreateToken(username, timespan)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	time.Sleep(timespan + time.Second)
	payload, err := maker.VerifyToken(token)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", payload)

	fmt.Println(time.Now().After(expiredAttime))
}
