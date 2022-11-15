package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Eric-lab-star/go-web-development/controllers"
	"github.com/Eric-lab-star/go-web-development/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	tpl, err := views.Parse(filepath.Join("template", "home.html"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHanlder(tpl))
	tpl, err = views.Parse(filepath.Join("template", "contact.html"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHanlder(tpl))
	tpl, err = views.Parse(filepath.Join("template", "faq.html"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHanlder(tpl))

	log.Println("server is listening to port http://localhost:3003")
	http.ListenAndServe("localhost:3003", r)
}
