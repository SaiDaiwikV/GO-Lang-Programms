package main

import (
	"fmt"
)

func printSlice[T any](items []T){
	for _, item := range items{
		fmt.Println(item)
	}
}

func printArray[T string | int | bool](items []T){ // we can use 'comparable' and 'interface{}' in [T any] in place of 'any'
	for _, item := range items{
		fmt.Println(item)
	}
}

// generics in structs too as follows:

type stack[T any] struct{
	elements []T
}



func main(){
	nums := []int{1,2,3,4,5}
	names := []string{"golang","python","cpp"}
	vals := []bool{true,false,false,true}
	printSlice(nums)
	printSlice(names)

	printArray(vals)
	printArray(names)

	myStack1 := stack[string]{
		elements: []string{"golang","cpp","python"},
	}

	fmt.Println(myStack1)

	myStack2 := stack[bool]{
		elements: []bool{false,true,false},
	}

	fmt.Println(myStack2)
}