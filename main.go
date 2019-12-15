package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: write "Hello, World!" using the http.ResponseWriter object
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
