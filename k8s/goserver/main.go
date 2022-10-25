package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello %s age %s", name, age)
	})

	http.HandleFunc("/secret", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "User %s password %s", user, password)
	})
	http.ListenAndServe(":8080", nil)
}
