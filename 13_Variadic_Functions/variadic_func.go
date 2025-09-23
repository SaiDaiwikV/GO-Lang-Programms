package main

import (
	"fmt"
)

func sum(nums ... int)int{
	total := 0

	for _, num := range nums{
		total = total + num
	}

	return total
}

func main(){

	// fmt.Println will accept 'n' number of parameters so it is a variadic function;
	fmt.Println(1,2,3,"golang",true)

	result := sum(4,5,6,7)

	fmt.Println(result)
}