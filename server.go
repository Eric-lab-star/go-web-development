package main

import (
	"log"
	"net/http"

	"github.com/Eric-lab-star/go-web-development/controllers"
	"github.com/Eric-lab-star/go-web-development/templates"
	"github.com/Eric-lab-star/go-web-development/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHanlder(
		views.Must(
			views.ParseFS(templates.FS, "layout.html", "home.html"),
		)))

	r.Get("/contact", controllers.StaticHanlder(
		views.Must(
			views.ParseFS(templates.FS, "layout.html", "contact.html"),
		)))

	r.Get("/faq", controllers.FAQ(
		views.Must(
			views.ParseFS(templates.FS, "layout.html", "faq.html"),
		)))

	log.Println("server is listening to port http://localhost:3003")
	http.ListenAndServe("localhost:3003", r)

}
