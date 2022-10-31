package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world</h1>")
}

func main() {
	http.HandleFunc("/", handleHome)
	fmt.Println("starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
