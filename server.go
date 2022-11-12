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
	r.Get("/contact", handleContact)
	log.Println("server is listening to port 3003")
	http.ListenAndServe("localhost:3003", r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("template", "home.html")
	exeTemplate(w, "", tplPath)

}

func handleContact(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["Name"] = "Kim"
	data["Email"] = "cyon256@naver.com"
	tplPath := filepath.Join("template", "contact.html")
	exeTemplate(w, data, tplPath)

}

func exeTemplate(w http.ResponseWriter, data any, tplPath string) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "Found Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
