package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/", Hello)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8080", nil)
}

func Hello(rw http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(rw, "Hello %s age %s", name, age)
}

func Secret(rw http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(rw, "User %s password %s", user, password)
}

func Healthz(rw http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() > 25 {
		rw.WriteHeader(500)
		rw.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		rw.WriteHeader(200)
		rw.Write([]byte("Ok"))
	}
}
