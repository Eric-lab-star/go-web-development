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

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering emails 24/7, though response times may be a bit slower on weekends",
		},
		{
			Question: "How do I contact support?",
			Answer:   "Email us cyon256@icloud.com",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}

func Contact(tpl views.Template) http.HandlerFunc {
	contact := struct {
		Name  string
		Email string
	}{
		Name:  "Kim",
		Email: "cyon256@icloud.com",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, contact)
	}
}
