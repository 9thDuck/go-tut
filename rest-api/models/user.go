package models

import (
	"regexp"
	"time"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
	"github.com/golang-jwt/jwt/v5"
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

func (u *User) VerifyUser(hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
}

func (u *User) ValidateCredentialsAndGetUser() (userPointer *User, err error) {
	userPointer, err = GetUserByEmail(u.Email)

	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPointer.Password), []byte(u.Password))

	if err != nil {
		return
	}

	userPointer.Password = ""

	return
}

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (u *User) IsEmailValid() bool {
	return emailRegex.MatchString(u.Email)
}

func GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, email, password FROM users WHERE email=?`

	row := db.Db.QueryRow(query, email)

	var user User = User{}

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (u *User) GenerateToken() (jwtStr string, err error) {
	JWT_SECRET, err := utils.GetEnvVariable("JWT_SECRET")

	if err != nil {
		return "", err
	}

	jwtDuration, err := utils.GetJwtDuration()

	if err != nil {
		return "", err
	}

	jwtPointer := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{"email": u.Email, "user_id": u.ID, "exp": time.Now().Add(jwtDuration).Unix()})

	jwtStr, err = jwtPointer.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return "", err
	}

	return jwtStr, nil
}
