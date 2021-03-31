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

	for index, _ := range users.Results {
		fmt.Println(users.FullName(index))
	}

	fmt.Println(users.Summary(1))
}
