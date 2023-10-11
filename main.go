package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	api "sesi-5/api"
	controller "sesi-5/controllers"

	model "sesi-5/models"
)

func init() {
	gob.Register(&model.User{})
}

func main() {
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/register", api.Register)
	http.HandleFunc("/login", controller.LoginPage)
	http.HandleFunc("/register", controller.RegisterPage)
	http.HandleFunc("/profile", controller.ProfilePage)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
