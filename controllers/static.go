package controllers

import (
	"net/http"

	"github.com/Eric-lab-star/go-web-development/views"
)

func StaticHanlder(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
