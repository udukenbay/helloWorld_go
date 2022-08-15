package main

import (
	"fmt"
	"hello_world/db"
)

func main()  {
	
	id, ivan := db.NewUser("Ivan", "Ivanov", "ivanov@gmail.com")
	db.SetUser(id, ivan)
	fmt.Println(db.GetUsers())

	id, user := db.NewUser("User", "Userov", "userov@gmail.com")
	db.SetUser(id, user)
	fmt.Println(db.GetUsers())

	db.RemoveUser(id)
	fmt.Println(db.GetUsers())
}