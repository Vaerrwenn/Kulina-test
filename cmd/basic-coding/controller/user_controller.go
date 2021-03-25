package controller

import (
	"bufio"
	"fmt"
	"kulina-test/cmd/basic-coding/entity"
	"os"
)

// UserRegistration handles User registration.
func UserRegistration() entity.User {
	var user entity.User
	scanner := bufio.NewScanner(os.Stdin)

	for isRunning := true; isRunning; {
		fmt.Print("Insert Name: ")
		scanner.Scan()
		user.Name = scanner.Text()

		if len(user.Name) < 4 {
			fmt.Println("Name must be more than 3 characters.")
		} else {
			isRunning = false
		}
	}

	return user
}

// UserInputAddress handles User address input.
func UserInputAddress(user *entity.User) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Address: ")
	scanner.Scan()
	user.Address = scanner.Text()
}
