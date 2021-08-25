package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	// controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("test@gmail.com")
	// fmt.Println(user)

	// session, _ := user.CreateSession()
	// fmt.Println(session)
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
	controllers.StartMainServer()

}
