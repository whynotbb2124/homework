package main

import (
	"io"
	"net/http"
	"strings"
)

var users []string

func main() {

	http.HandleFunc("/hello", hellofunc)
	http.HandleFunc("/registration", registerfunc)
	http.HandleFunc("/users", usershandler)
	http.HandleFunc("/", welcomePage)

	println("host started on :8080")

	http.ListenAndServe(":8080", nil)
}

func hellofunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, you are in main page"))
}

func registerfunc(w http.ResponseWriter, r *http.Request) {

	user, _ := io.ReadAll(r.Body)

	users = append(users, string(user))

	w.Write([]byte("User added"))
}

func usershandler(w http.ResponseWriter, r *http.Request) {

	usersbase := strings.Join(users, ", ")

	w.Write([]byte("Users list: " + usersbase))
}
