package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<a href=\"/contact\"><h1>contact page</h1></a>")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:cyon256@icloud.com\">cyon256@icloud.com</a></p>")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {

	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(homeHandler))
	if err != nil {
		fmt.Println("err", err)
	}
}
