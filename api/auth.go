package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	model "sesi-5/models"

	sessionStore "sesi-5/config"

	"github.com/gorilla/sessions"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Job      string `json:"job"`
	Reason   string `json:"reason"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseMessageWithUserData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	User    *model.User `json:"user,omitempty`
}

var store = sessions.NewCookieStore([]byte("sesi-5-go"))

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error decoding request:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, ok := model.Users[req.Email]
	if !ok || !model.PasswordCheck(user.Password, req.Password) {
		response := ResponseMessage{
			Status:  "failed",
			Message: "Invalid email or password",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	session, err := sessionStore.Store.Get(r, "user-session")
	if err != nil {
		fmt.Println("Error retrieving session:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	userDetails := model.GetUserDetails(req.Email)
	session.Values["userDetails"] = userDetails
	err = session.Save(r, w)
	if err != nil {
		fmt.Println("Error saving session:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Stored session values:", session.Values)

	response := ResponseMessageWithUserData{
		Status:  "success",
		Message: "Logged in successfully",
		User:    userDetails,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, exists := model.Users[req.Email]
	if exists {
		response := ResponseMessage{
			Status:  "failed",
			Message: "User already exists",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	model.AddUser(req.Email, req.Password, req.Address, req.Job, req.Reason)

	response := ResponseMessageWithUserData{
		Status:  "success",
		Message: "User successfully registered",
		User:    model.GetUserDetails(req.Email),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
