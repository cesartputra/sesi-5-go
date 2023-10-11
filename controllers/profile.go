package controller

import (
	"html/template"
	"net/http"
	sessionStore "sesi-5/config"
	model "sesi-5/models"
)

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Store.Get(r, "user-session")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	userDetailsInterface, ok := session.Values["userDetails"]
	if !ok {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	userDetails, ok := userDetailsInterface.(*model.User)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Email   string
		Address string
		Job     string
		Reason  string
	}{
		Email:   userDetails.Email,
		Address: userDetails.Address,
		Job:     userDetails.Job,
		Reason:  userDetails.Reason,
	}

	tmpl, err := template.New("profile").ParseFiles("views/profile.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "profile.html", data)
}
