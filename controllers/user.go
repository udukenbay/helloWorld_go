package controllers

import (
	"encoding/json"
	"fmt"
	"hello_world/models"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET users")

	users := models.GetUsers()
	data, err := json.Marshal(users)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, string(data))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := models.GetUser(vars["id"])
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, string(data))
}

// POST user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("POST user")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	defer r.Body.Close()

	user := models.User{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	models.SetUser(user)
	fmt.Fprint(w, "DONE")
}

// DELETE user/{id}
func RemoveUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("DELETE user")

	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		models.RemoveUser(id)
		fmt.Fprint(w, "DONE")
		return
	}
	fmt.Fprint(w, "Необходимо указать id пользователя")
}
