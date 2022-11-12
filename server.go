package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	var r = chi.NewRouter()
	r.Get("/", handleHome)
	http.ListenAndServe("localhost:3000", r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}
