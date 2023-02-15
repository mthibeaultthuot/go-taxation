package main

import (
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "NewPassword1234!"
	hashPassword, err := HashPassword(password)
	if err != nil {
		return
	}
	err = CheckPasswordHash(password, hashPassword)
	if err != nil {
		t.Fatalf("password don't match")
	}
}
