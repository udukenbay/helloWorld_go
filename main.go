package main

import (
	"fmt"
	"hello_world/controllers"
	"hello_world/models"
	"hello_world/src/server"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	// GET user
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	// GET user/{id}
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	// POST user
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	// DELETE user/{id}
	router.HandleFunc("/user/{id}", controllers.RemoveUser).Methods("DELETE")

	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		t, _ := template.ParseFiles("views/home.html")

		users := models.GetUsers()
		t.Execute(w, users)
	})

	fmt.Println("Веб сервер запущен")
	s := server.New(7878)
	s.Start(router)
}
