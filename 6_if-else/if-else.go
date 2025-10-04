package main

import "fmt"

func main(){
	var age int = 18

	if age >= 18{
		fmt.Println("You are eligible to vote in india.")
	}else if age < 0{
		fmt.Println("Please enter the valid age number to run the program.")
	}else{
		fmt.Println("You are not eligible to vote in india.")
	}


	// ------------------------------------------------------------- //

	role := "admin"
	hasPermissions := false

	if role == "admin" || hasPermissions == true {
		fmt.Println("Your are at the admin role for this application.")
	}else{
		fmt.Println("Your are not at the admin level for this application.")
	}

	// -------------------------------------------------------------//

	if age := 22; age >= 18{
		fmt.Println("The person is adult.")
	}else if age > 12{
		fmt.Println("The person is a teenager.")
	}else {
		fmt.Println("The person is a kid.")
	}

}