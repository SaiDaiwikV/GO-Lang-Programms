package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(i int,wg *sync.WaitGroup){

	defer wg.Done()

	fmt.Println("doing task id",i)
}

func main(){

	var wg sync.WaitGroup
	
	for i := 0;i<=10;i++{
		wg.Add(1)
		go task(i,&wg)

		// func (i int){
		// 	fmt.Println(i)
		// }(i)
	}

	// time.Sleep(time.Second * 2)

	wg.Wait()

}