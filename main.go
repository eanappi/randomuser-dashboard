package main

import (
	"fmt"

	"github.com/eanappi/randomuser"
)

func main() {
	fmt.Println("Test")
	randomuser.Test()
	users, err := randomuser.NewRandomUserJson(10)
	if err != nil {
		fmt.Errorf("The service is offline by: %s", err)
	}

	fmt.Println(users)
}
