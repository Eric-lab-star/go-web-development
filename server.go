package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Eric-lab-star/go-web-development/controllers"
	"github.com/Eric-lab-star/go-web-development/templates"
	"github.com/Eric-lab-star/go-web-development/views"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.StaticHanlder(
		views.Must(
			views.ParseFS(templates.FS, "layout.html", "home.html"),
		)))

	r.Get("/contact", controllers.Contact(
		views.Must(
			views.ParseFS(templates.FS, "layout.html", "contact.html"),
		)))

	r.Get("/faq", controllers.FAQ(
		views.Must(
			views.ParseFS(templates.FS, "layout.html", "faq.html"),
		)))
	r.Get("/signup", controllers.StaticHanlder(
		views.Must(
			views.ParseFS(templates.FS, "layout.html", "signup.html"),
		)))
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "dist"))
	FileServer(r, "/files/", filesDir)
	log.Println("server is listening to port http://localhost:3003")
	http.ListenAndServe("localhost:3003", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())

		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
