package main

import (
	"fmt"
)

// it is passing by value
func changeNum(num int) {
	num = 5

	fmt.Println("In changeNum num value is : ",num)
}

func changeNump(num *int){
	*num = 5 // de-referrencing the address

	fmt.Println("In changeNump num value is : ", *num)

}

func main(){
	num := 1

	changeNum(num)

	fmt.Println("The Memory Address is ",&num)
	fmt.Println(num)

	changeNump(&num)
	fmt.Println(num)
}