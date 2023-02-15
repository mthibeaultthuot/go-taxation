package main

import (
	"context"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
	"log"
	"net/http"
)

func main() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})
	srv.SetUserAuthorizationHandler(test)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(r.Context(), w, r)
		if err != nil {
			log.Panic(err)
		}
		store.Set("LoggedInUserID", r.Form.Get("username"))
		err = store.Save()
		if err != nil {
			log.Panic(err)
		}
		log.Println("hello")
		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}

func test(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		log.Panic(err)
	}
	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
