package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Page</h1>")
}

func main() {
	// Setup some routes with http object
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3001", nil)

}
