package main

import (
	"fmt"
)

func add(a int,b int)int{
	return a + b
}

func getlanguages() (string, string, string){
	return "golang", "python", "cpp"
}

func processIt() func(a int) int{
	return  func(a int) int{
		return a
	}
}

func main(){
	res := add(2, 3)

	fmt.Println(res)

	lang1 , lang2 , lang3 := getlanguages()

	fmt.Println(lang1, lang2, lang3)

	// or you can simply call this function in fmt.Println directly
	
	fmt.Println(getlanguages())

	lang1 , lang2 , _ = getlanguages() // _ means the compiler will ignore that parameter.

	fmt.Println(lang1, lang2)

	fn := processIt()
	fmt.Println(fn(6))
	fmt.Println(fn(100))
}