package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("<h1>Web Server Go V2</h1>"))
	})
	http.ListenAndServe(":8080", nil)
}
