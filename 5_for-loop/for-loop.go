package main

import "fmt"

func main(){

	i := 1
	// this is will work as for loop:
	for i <= 5{
		fmt.Println(i)
		i++
	}

	// for infinte loop :

	// for{
	// 	fmt.Println(1)
	// }

	// classic for loop:
	for i := 0;i<=5;i++{
		if i == 4{
			break
		}
		if i == 2{
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("----------------")
	for i:= range 3{
		fmt.Println(i)
	}
}