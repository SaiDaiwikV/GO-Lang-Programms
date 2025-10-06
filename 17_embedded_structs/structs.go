package main

import (
	"fmt"
)

type customer struct{
	name string
	phone string
}

type order struct{
	id string
	amount float32
	status string
	customer      // this is import customer struct inside the order struct	
}


func main(){

	newCustomer := customer{
		name : "Daiwik",
		phone: "1234567890",
	}

	newOrder := order{
		id:"1",
		amount: 666.9,
		status: "Order Confirmed",
		customer: newCustomer,
	}

	fmt.Println(newOrder)

	newOrder.customer.name = "Robin"

	fmt.Println(newOrder)  // name has been changed

}