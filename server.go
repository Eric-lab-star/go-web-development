package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	w.Header().Add("transfer-encoding", "chunked")
	w.Header().Add("trailer", "transfer-encoding")
	w.Write([]byte("heelo"))
	fmt.Fprint(w, "<h1>Hi air is watching</h1>")
}

func main() {
	http.HandleFunc("/", handleHome)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("err", err)
	}
}
