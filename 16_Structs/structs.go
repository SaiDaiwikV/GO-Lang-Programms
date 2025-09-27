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

func (o *order) changeStatus(status string) { // the syntax is we need to use first letter of struct for example i used 'o' because my struct name is 'order'
	o.status = status
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

}
