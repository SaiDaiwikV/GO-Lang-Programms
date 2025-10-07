package main

import(
	"fmt"
)

type OrderStatus int

const(
	Received OrderStatus = iota
	Prepared
	Delivered
	Cancelled
	Shipped
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("chaging the order status to :", status)
}

func main(){
	// fmt.Println()
	changeOrderStatus(Received) 
	changeOrderStatus(Delivered)
	changeOrderStatus(Shipped)

	// this is how enums works.
}