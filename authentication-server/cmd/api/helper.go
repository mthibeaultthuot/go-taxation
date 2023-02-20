package main

import (
	"github.com/D3xt3rrrr/go-cloud/authentication-server/data"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var (
	Secret = "$asdf$asdft4$"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Secret+password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password string, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(Secret+password))
	if err != nil {
		log.Panic(err)
	}
	return err
}

func IsUserAdmin(user data.User) {

}
