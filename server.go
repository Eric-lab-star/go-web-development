package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hi  is watching</h1>")
}

func main() {
	http.HandleFunc("/", handleHome)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("err", err)
	}
}
