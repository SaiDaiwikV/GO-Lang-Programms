package main

import "fmt"

func main(){
	// var name string = "SaiDaiwik"
	// golang will "infer" the variable automatically

	var name = "SaiDaiwik"
	var isAdult = true
	var age = 22
	var price = 50.5


	fmt.Println("My name is " + name)
	fmt.Println(isAdult)
	fmt.Println(age)
	fmt.Println(price)

	// shorthand declrations:
	// colon is very important for short hand notations in variable decleration

	NAME := "SaiDaiwik"
	ISADULT := true
	AGE := 22
	PRICE := 50.5

	fmt.Println(NAME)
	fmt.Println(ISADULT)
	fmt.Println(AGE)
	fmt.Println(PRICE)


}