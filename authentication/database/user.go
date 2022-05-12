package database

import (
	"errors"
	"log"

	"github.com/uptrace/bun"
)

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type user struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID           int64 `bun:",pk,autoincrement"`
	Email        string
	Username     string
	PasswordHash string
	Fullname     string
	CTimeStamp   string
	UTimeStamp   string
	Role         string
}

func GetUser(email, password string) (*user, bool) {
	usr, err := getUser(email)
	if err != nil {
		return nil, false
	}

	passwordHash, err := generatehashPassword(password)
	if err != nil {
		log.Println("error in password hash")
		return nil, false
	}

	if !usr.validatePassword(passwordHash) {
		return nil, false
	}

	return usr, true
}

// Check if the passwordhash is valid
func (u *user) validatePassword(passwordHash string) bool {
	return u.PasswordHash == passwordHash
}

func AddUser(Email, Username, Password, Fullname, Role string) error {
	u, _ := getUser(Email)
	if u != nil {
		return errors.New("user already exists")
	}
	PasswordHash, err := generatehashPassword(Password)
	if err != nil {
		log.Println("error in password hash")
		return err
	}
	return insertUser(Email, Username, PasswordHash, Fullname, Role)
}
