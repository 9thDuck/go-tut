package models

import (
	"regexp"

	"example.com/rest-api/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO USERS (email, password) VALUES(?, ?)`

	statement, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)

	if err != nil {
		return err
	}

	defer statement.Close()

	userInsertRes, err := statement.Exec(u.Email, string(hashedPasswordBytes))

	if err != nil {
		return err
	}

	userId, err := userInsertRes.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = userId

	return nil
}

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (u *User) IsEmailValid() bool {
	return emailRegex.MatchString(u.Email)
}
