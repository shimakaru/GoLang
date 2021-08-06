package main

import (
	"fmt"
)

type user struct {
	Name string
	Age  int
}
type T struct {
	user
}

func (u *user) SetName() {
	u.Name = "A"
}

type users []*user

func NewUser(Name string, Age int) *user {
	return &user{Name: Name, Age: Age}
}
func main() {
	user1 := user{Name: "A", Age: 10}
	users := users{}
	users = append(users, &user1)
	for i, u := range users {
		fmt.Println(i, *u)
	}

}
