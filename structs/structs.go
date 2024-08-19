package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println("name:", u.firstName, u.lastName)
	fmt.Println("Born on", u.birthDate)
	fmt.Println("Member since", u.createdAt)
}

func (u *user) resetUserData() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}
	return &user{firstName: firstName, lastName: lastName, birthDate: birthDate, createdAt: time.Now()}, nil
}

func main() {
	firstName := getUserInputData("Please enter your first name:")
	lastName := getUserInputData("Please enter your last name:")
	birthDate := getUserInputData("Please enter your birthdate(DD/MM/YYYY):")

	userData, err := newUser(
		firstName, lastName, birthDate,
	)

	if err != nil {
		panic(err)
	}

	userData.outputUserDetails()
	userData.resetUserData()
	userData.outputUserDetails()

}

func getUserInputData(label string) (userInput string) {
	fmt.Print(label)
	fmt.Scanln(&userInput)
	return userInput
}
