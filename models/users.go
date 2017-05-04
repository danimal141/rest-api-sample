package models

import "fmt"

type User struct {
	Id       int
	UserName string
}

func SampleUsers() []User {
	users := make([]User, 0, 3)
	for i := 0; i < 3; i++ {
		users = append(users, User{Id: i, UserName: fmt.Sprint("testuser", i)})
	}
	return users
}
