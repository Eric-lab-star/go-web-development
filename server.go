package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	var r = chi.NewRouter()
	r.Get("/", handleHome)
	log.Println("server is listening to port 3003")
	http.ListenAndServe("localhost:3003", r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tplPath := filepath.Join("template", "hello.html")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("parsing template: %v", err)

		return
	}
	err = tpl.Execute(w, "string")
	if err != nil {
		log.Printf("executing template: %v", err)

		return
	}
}
