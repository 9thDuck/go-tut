package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	createdAt := time.Now()

	return &User{firstName, lastName, birthDate, createdAt}, nil
}

func (u *User) OutputUserDetails() {
	fmt.Println("name:", u.firstName, u.lastName)
	fmt.Println("Born on", u.birthDate)
	fmt.Println("Member since", u.createdAt)
}

func (u *User) ResetUserData() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

// admin
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email string, password string, user *User) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	return &Admin{
		email: email, password: password, User: *user,
	}, nil
}
