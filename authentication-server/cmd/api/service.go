package main

import (
	"encoding/json"
	"github.com/D3xt3rrrr/go-cloud/authentication-server/data"
	"log"
	"net/http"
)

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := data.FetchUser(username)
	if err != nil {
		log.Panic(err)
	}
	err = CheckPasswordHash(password, user.Password)
	if err != nil {
		log.Panic(err)
	}
	bearer := "Bearer " + GenerateJwt(user)
	w.Header().Set("Authorization", bearer)
	log.Println("authorization : ", w.Header().Get("Authorization"))
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
}

func (app *Config) Registration(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	hashPassword, err := HashPassword(password)
	if err != nil {
		log.Panic(err)
	}
	_, err = data.FetchUser(username)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode("user already exist")
		return
	}
	user := data.User{
		Username:   username,
		Password:   hashPassword,
		Permission: 1,
	}
	err = data.InsertUser(user)
	if err != nil {
		log.Panic(err)
	}
	log.Println("new registration ==> ", user.Username)
}

func (app *Config) Delete(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := data.FetchUser(username)
	if err != nil {
		log.Panic(err)
		return
	}
	err = CheckPasswordHash(password, user.Password)
	if err != nil {
		log.Panic(err)
		return
	}

	err = data.DeleteUser(user)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("user deleted")
}

func (app *Config) Refresh(w http.ResponseWriter, r *http.Request) {

}

func (app *Config) Token(w http.ResponseWriter, r *http.Request) {

}
