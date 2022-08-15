package db

import "github.com/google/uuid"

type User struct {
	Id int
	Firstname string
	Lastname string
	Email string
}

type Users map[int]User

var users = Users{}

func (u *User) set(id string, firstname, lastname string, email string) {
	u.Id = id
	u.Firstname = firstname
	u.Lastname = lastname
	u.Email = email
}

func NewUser(firstname, lastname string, email string) (id int, user User) {
	id = newUuid
	user.set(id, firstname, lastname, email)
	return 
}

//read data/get data from map
func GetUser(id string) User {
	return users[id]
}

func GetUsers() (u []User) {
	for _, value := range users {
		u = append(u, value)
	}
	return u
}

//create/add data to map
//update
func SetUser(id string, user User) {
	users[id] = user
}

//delete data from map
func RemoveUser(id string) {
	delete(users, id)
}

func newUuid() string {
	id, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return id.String()
}