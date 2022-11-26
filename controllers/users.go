package controllers

import (
	"net/http"

	"github.com/Eric-lab-star/go-web-development/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// we need a view to render
	u.Templates.New.Execute(w, nil)
}
