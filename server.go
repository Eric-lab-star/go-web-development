package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	var r = chi.NewRouter()
	r.Get("/", handleHome)
	r.Get("/contact", handleContact)
	r.Get("/faq", handleFaq)
	log.Println("server is listening to port 3003")
	http.ListenAndServe("localhost:3003", r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("running home")
	tplPath := filepath.Join("template", "home.html")
	exeTemplate(w, "", tplPath)

}

func handleContact(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Name":  "Kim",
		"Email": "cyon256@naver.com",
	}
	tplPath := filepath.Join("template", "contact.html")
	exeTemplate(w, data, tplPath)

}

func handleFaq(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"Name": "Kim"}
	exeTemplate(w, data, filepath.Join("template", "faq.html"))
}

func exeTemplate(w http.ResponseWriter, data any, tplPath string) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "Found Error parsing template", http.StatusInternalServerError)
		return
	}
	fmt.Println("hello")
	log.Println("log print")
	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
