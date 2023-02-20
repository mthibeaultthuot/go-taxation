package main

import (
	"fmt"
	"github.com/D3xt3rrrr/go-cloud/authentication-server/data"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth
var secret = "Adfsafd#W#Rasdfasf$"

func GenerateJWTToken(user data.User) string {
	tokenAuth := jwtauth.New("HS256", []byte(secret+user.Password), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": user.Username})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	return tokenString
}
