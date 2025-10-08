package main

import (
	"fmt"
	// "time"
	// "time"
	// "math/rand"
	// "time"
)

// func processNum(numChan chan int){

// 	for num := range numChan{
// 	fmt.Println("Processing number ",num)
// 	time.Sleep(time.Second)
// 	}
// }

// func task(done chan bool){
// 	defer func(){done <- true}()

// 	fmt.Println("Processing....")

// }

// func sum(result chan int,num1 int,num2 int){
// 	numRes := num1 + num2;
// 	result <- numRes
// }

// func emailSender(emailChan chan string,done chan bool){
// 	defer func() {done <- true}()

// 	for email := range emailChan{
// 		fmt.Println("Sending email to",email)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func main(){

	// chanNum := make(chan int)

	// go sum(chanNum,50,100)

	// res := <-chanNum

	// fmt.Println(res)

	// done := make(chan bool)

	// go task(done)
	 
	// <-done

	// emailChan := make(chan string, 100) // giving bufer size of 100;

	// go emailSender(emailChan, done)

	// for i := 0;i<100;i++{
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("Completed Sending.")

	// close(emailChan)
	// <-done


	// go processNum(chanNum)

	// chanNum <- 5

	// for{
	// 	chanNum <- rand.Intn(100)
	// }

	chan1 := make(chan int)
	chan2 := make(chan string)


	go func(){
		chan1 <- 10
	}()

	go func(){
		chan2 <- "ping"
	}()

	for i:= 0;i<2;i++{
		select {
		case chan1Val := <-chan1:
			fmt.Println("Recevied data from chan1 that is",chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Recevied data from chan2 that is",chan2Val)
		}
	}

}
