package main

import (
	"fmt"
	// "slices"
)

func main(){
	// creating a map

	m := make(map[string]string)

	// setting element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])
	// if key does not exist in the map then it returns zero value;
	fmt.Println(m["phone"])


	m1 := make(map[string]int)

	m1["CPP"] = 1
	m1["GOLANG"] = 2
	m1["PYTHON"] = 3

	fmt.Println(m1)

	delete(m1,"CPP")

	fmt.Println(m1)

}