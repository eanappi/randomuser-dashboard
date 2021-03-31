package main

import (
	"fmt"

	"github.com/eanappi/randomuser"
)

func main() {
	users, err := randomuser.NewUsers(10)
	if err != nil {
		fmt.Errorf("The service is offline by: %s", err)
	}

	user := users.Results

	fmt.Println(user[0].Gender)
}
