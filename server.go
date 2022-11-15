package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Eric-lab-star/go-web-development/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	var r = chi.NewRouter()
	r.Get("/", handleHome)
	r.Get("/contact", handleContact)
	r.Get("/faq", handleFaq)
	log.Println("server is listening to port http://localhost:3003")
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
	data := map[string]string{"Name": ""}
	exeTemplate(w, data, filepath.Join("template", "faq.html"))
}

func exeTemplate(w http.ResponseWriter, data any, filepath string) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := views.Parse(filepath)

	if err != nil {
		http.Error(w, "Found Error parsing template", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)

}
