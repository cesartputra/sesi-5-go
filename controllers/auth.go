package controller

import (
	"html/template"
	"net/http"
	model "sesi-5/models"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("login").ParseFiles("views/login.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "login.html", struct {
		ErrorMessage string
		Emails       []string
	}{
		ErrorMessage: "", Emails: model.GetUsersEmail(),
	})
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("register").ParseFiles("views/register.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "register.html", nil)
}
