package main

import (
	"fmt"
	"os"

	"goprojects/users"
)

func main() {
	username, password := os.Getenv("GMAIL_USERNAME"), "qwerty123"

	err := users.NewUser(username, password)
	if err != nil {
		fmt.Printf("Couldn't create user: %s\n", err.Error())
		return
	}

	err = users.AuthenticateUser(username, password)
	if err != nil {
		fmt.Printf("Couldn't authenticate user: %s\n", err.Error())
		return
	}
}
