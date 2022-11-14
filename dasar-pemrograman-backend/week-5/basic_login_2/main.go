package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/login", loginHandler)
	http.Handle("/dashboard", dashboardHandler)
	http.ListenAndServe("localhost:8080", nil)
}

var (
	loginHandler     = MiddlewarePost(HandleLogin())
	dashboardHandler = MiddlewareGet(MiddlewareAuth(HandleDashboard()))
)

type User struct {
	Username string
	Password string
}

var users []User = []User{{"user1", "passuser1"}, {"user2", "passuser2"}}

func HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error form parsing"))
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		token, err := getAuthToken(username, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	}
}

func HandleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Dashboard!"))
	}
}

func MiddlewarePost(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("HTTP method should be POST"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func MiddlewareGet(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("HTTP method should be GET"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func MiddlewareAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authorization header not found"))
			return
		}
		user := getUser(username, password)
		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("user not found or incorrect password"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func getUser(username, password string) *User {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user
		}
	}
	return nil
}

func getAuthToken(username, password string) (string, error) {
	user := getUser(username, password)
	if user != nil {
		return generateAuthToken(username, password), nil
	}
	return "", errors.New("user not found or incorrect password")
}

func generateAuthToken(username, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", username, password)))
}
