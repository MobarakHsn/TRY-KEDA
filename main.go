package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the Go HTTP server!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
