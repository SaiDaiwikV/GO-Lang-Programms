package main

import (
	"fmt"
	"time"
)

func task(i int){
	fmt.Println("doing task id",i)
}

func main(){
	
	for i := 0;i<=10;i++{
		go task(i)

		func (i int){
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)

}