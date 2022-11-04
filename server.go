package main

import (
	"fmt"
	"net/http"
)

func homHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<a href=\"/contact\"><h1>contact page</h1></a>")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:cyon256@icloud.com\">cyon256@icloud.com</a></p>")
}

func main() {
	http.HandleFunc("/", homHandler)
	http.HandleFunc("/contact", contactHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("err", err)
	}
}
