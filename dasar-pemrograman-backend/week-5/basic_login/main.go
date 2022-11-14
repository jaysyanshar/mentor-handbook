package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
)

const host = "localhost:8080"

type User struct {
	Username string
	Password string
}

var data []User = []User{{"username1", "password1"}, {"username2", "password2"}}

func main() {
	http.Handle("/login", HandlePostMiddleware(HandleLogin()))
	http.Handle("/dashboard", HandleDashboard())
	fmt.Println("listening to " + host)
	http.ListenAndServe(host, nil)
}

func HandlePostMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func HandleGetMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New("Method not allowed")
	}
	return nil
}

func HandleAuthMiddleware(w http.ResponseWriter, r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("unauthorized")
	}
	user := getUser(username, password)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		errors.New("unauthorized")
	}
	return nil
}

func HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("form error"))
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user := getUser(username, password)
		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Basic " +
			base64.StdEncoding.EncodeToString([]byte(user.Username+":"+user.Password))))
	}
}

func HandleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := HandleGetMiddleware(w, r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		err = HandleAuthMiddleware(w, r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	}
}

func getUser(username, password string) *User {
	for _, user := range data {
		if user.Username == username && user.Password == password {
			return &user
		}
	}
	return nil
}
