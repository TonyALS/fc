package main

import ("net/http"
		"os"
		"fmt"
	)

func main() {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello %s age %s", name, age)
	})
	http.ListenAndServe(":8080", nil)
}
