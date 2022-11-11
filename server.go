package main

import (
	"errors"

	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad request"))
	}
}

func main() {
	r := chi.NewRouter()
	r.Method("GET", "/", Handler(customHandler))
	r.Method("GET", "/movies", Handler(movieHandler))
	http.ListenAndServe("localhost:8080", r)
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")
	if q != "" {
		return errors.New(q)
	}
	w.Write([]byte("foo"))
	return nil
}

func movieHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")
	if q != "" {
		return errors.New(q)
	}
	w.Write([]byte("movies"))
	return nil
}
