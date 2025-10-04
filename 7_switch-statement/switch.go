package main

import (
	"fmt"
	"time"
)

func main(){
	// simple switch statement:
	i := 4

	switch(i){
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Other number")
	}

	// multiple switch statement :
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend.")
	default:
		fmt.Println("It's a Working Day.")
	}

	// type switch
	WhoAmI := func(i interface{}){
		switch j := i.(type){
		case int:
			fmt.Println("It's an INTEGER")
		case string:
			fmt.Println("It's a STRING")
		case bool:
			fmt.Println("It's a BOOL")
		default:
			fmt.Println("OTHER",j)
		}	
	}

	WhoAmI("golang")
}