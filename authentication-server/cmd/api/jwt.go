package main

import (
	"github.com/D3xt3rrrr/go-cloud/authentication-server/data"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
)

var secret = []byte("asdfadsfdasfadsfd")

type JwtClaims struct {
	Username   string `json:"username"`
	Permission int    `json:"permission"`
	jwt.RegisteredClaims
}

func GenerateJwt(user data.User) string {
	expiration := time.Now().Add(5 * time.Minute)
	claims := &JwtClaims{
		Username:   user.Username,
		Permission: user.Permission,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
	// TODO : add security with jwt.SigningMethodECDSA
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	return jwtToken
}

func VerifyJwt(JwtToken string) (bool, data.User) {
	// TODO : add more security and verification
	log.Println(JwtToken)
	token, err := jwt.ParseWithClaims(JwtToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, data.User{}
	}
	claims := token.Claims.(*JwtClaims)
	return token.Valid, data.User{Username: claims.Username, Permission: claims.Permission}
}

func getTokenFromBearer(bearerToken string) string {
	return strings.Split(bearerToken, "Bearer ")[1]
}
