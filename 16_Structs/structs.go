package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

// if i want to create a function that changes the value of status then;
// i need to create a link between function and struct
// receiver type:

func newOrder(id string, amount float32, status string) *order {
	Myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &Myorder
}

func (o *order) changeStatus(status string) { // the syntax is we need to use first letter of struct for example i used 'o' because my struct name is 'order'
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {
	myorder1 := order{
		id:     "1",
		amount: 360,
		status: "Processing",
	}

	myorder1.createdAt = time.Now()

	fmt.Println(myorder1.amount)

	fmt.Println(myorder1)

	fmt.Println("----------------------------")

	myorder2 := order{
		id:        "2",
		amount:    999.9,
		status:    "Payemnt Pending",
		createdAt: time.Now(),
	}

	// i can update myorder1 details here too :
	myorder1.status = "Paid"

	fmt.Println("Details of MyOrder1 : ", myorder1)
	fmt.Println("Details of MyOrder2 : ", myorder2)

	fmt.Println("----------------------------")

	// using changeStatus function to update myorder2:

	myorder2.changeStatus("Payment Failed")

	fmt.Println("Updated myorder2 : ", myorder2)

	// get the amount value from functions which is linked with struct:

	fmt.Println("The amount of order-1 is : ", myorder1.getAmount())

	fmt.Println("The amount of order-2 is :  ", myorder2.getAmount())

	fmt.Println("-----------------------------")

	// creating myorder3 using constructor of struct:
	myorder3 := newOrder("3", 666.99, "Order Initiated")

	fmt.Println("Order-3 is : ", *myorder3)

	// in-line structs:

	lang := struct {
		name   string
		isGood bool
	}{"Golang", true}

	fmt.Println(lang)
}
