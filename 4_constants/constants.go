package main

import "fmt"

const age = 22 // this can be written above the main function

func main(){

	const name = "SaiDaiwik"

	const(
		host = "localhost"
		port = 8080
	) // this is used for multiple variables to be constant

	fmt.Println(name, age)
	fmt.Println(host, port)
}