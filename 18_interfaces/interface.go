package main

import (
	"fmt"
)

type paymenter interface{
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) MakingPayment(amount float32){
	// razopayGW := razopay{}
	// razopayGW.pay(amount)

	// stripeGW := stripe{}
	p.gateway.pay(amount)

}

type razopay struct {}

func (r razopay) pay(amount float32){
	fmt.Println("Making payment of amount",amount, "using razopay")
}

type stripe struct {}

func (s stripe) pay(amount float32){
	fmt.Println("Making payment of amount", amount ,"using stripe")
}

type fakepayment struct {}

func (f fakepayment) pay(amount float32){
	fmt.Println("Making payment of amount ", amount , "using fakepayment")
}

func main(){
	// stripeGW := stripe{}
	razoPaymentGW := razopay{}
	stripeGW := stripe{}
	fakeGW := fakepayment{}
	Mypayment := payment{
		gateway: razoPaymentGW,
	}

	Mypayment.MakingPayment(100)

	Mypayment = payment{
		gateway: stripeGW,
	}

	Mypayment.MakingPayment(215)

	Mypayment = payment{
		gateway: fakeGW,
	}

	Mypayment.MakingPayment(1000)

}