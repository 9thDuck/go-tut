package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserInputData("Please enter your first name:")
	lastName := getUserInputData("Please enter your last name:")
	birthDate := getUserInputData("Please enter your birthdate(DD/MM/YYYY):")

	userData, err := user.New(
		firstName, lastName, birthDate,
	)

	if err != nil {
		panic(err)
	}

	userData.OutputUserDetails()
	userData.ResetUserData()
	userData.OutputUserDetails()
}

func getUserInputData(label string) (userInput string) {
	fmt.Print(label)
	fmt.Scanln(&userInput)
	return userInput
}
