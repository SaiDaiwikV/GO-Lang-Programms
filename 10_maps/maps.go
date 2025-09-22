package main

import (
	"fmt"
	"maps"
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

	clear(m1)

	fmt.Println(m1)

	m2 := map[string]string{"phone":"Poco F7", "Price":"29k"}

	fmt.Println(m2)

	v1, ok := m2["Price"]
	v2, kk := m2["mobile"]
	if ok {
		fmt.Println("All Okay")
	}else{
		fmt.Println("Not Okay")
	}

	if kk { // kk is not present in m2 map
		fmt.Println("Phone is Okay")
	}else{
		fmt.Println("Phone is not Okay")
	}

	// printing v1 and v2 values :
	fmt.Println(v1)
	fmt.Println(v2)

	// comparing two maps in two different ways:

	n1 := map[string]string{"name":"goalng", "IDE":"VS-CODE"}
	n2 := map[string]string{"name":"goalng", "IDE":"VS-CODE"}
	

	fmt.Println(maps.Equal(n1,n2))

	n1 = map[string]string{"name":"goalng", "IDE":"VS-CODE"}
	n2 = map[string]string{"name":"CPP", "IDE":"VS-CODE"}

	fmt.Println(maps.Equal(n1,n2))
	
}