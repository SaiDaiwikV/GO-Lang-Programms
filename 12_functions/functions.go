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

func main(){
	res := add(2, 3)

	fmt.Println(res)

	lang1 , lang2 , lang3 := getlanguages()

	fmt.Println(lang1, lang2, lang3)

	// or you can simply call this function in fmt.Println directly
	
	fmt.Println(getlanguages())
}