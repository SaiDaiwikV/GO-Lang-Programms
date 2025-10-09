package main

import (
	"fmt"
	"myPackages/auth"
	"myPackages/user"

	"github.com/fatih/color"
	// "github.com/labstack/gommon/color"
)

func main(){
	auth.LoginwithCredentials("go-lang","secretpassword")

	sessionOutput := auth.GetSession() // returns string so saving to a variable to print in console.

	fmt.Println(sessionOutput)

	user1 := user.UserDetails{
		Email: "go-lang@gmail.com",
		Name: "Go-lang",
	}

	fmt.Println("Details of user-1 is :",user1)
	color.Red(user1.Email)
}
