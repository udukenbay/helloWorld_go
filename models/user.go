package models

import "github.com/google/uuid"

type User struct {
	Id        string `json:"id" xml:"ID"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string
}

type Users map[string]User

var users = Users{}

func (u *User) set(id string, firstname, lastname string, email string) {
	u.Id = id
	u.Firstname = firstname
	u.Lastname = lastname
	u.Email = email
}

func NewUser(firstname, lastname string, email string) (id string, user User) {
	id = NewUUID() //функция для формирования id
	user.set(id, firstname, lastname, email)
	return
}

//read-get data from map
func GetUser(id string) *User {
	if value, ok := users[id]; ok {
		return &value
	}
	return nil
}

//read slice data from map
func GetUsers() (u []User) {
	for _, value := range users {
		u = append(u, value)
	}
	return
}

//create/add data to map
//update
func SetUser(user User) {
	id := NewUUID()
	user.Id = id
	users[id] = user
}

//delete data from map
func RemoveUser(id string) {
	delete(users, id)
}

func NewUUID() string {
	return uuid.New().String()
}
