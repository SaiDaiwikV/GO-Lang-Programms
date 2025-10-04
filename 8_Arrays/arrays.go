package main

import "fmt"

func main(){

	// creating an array of any data types will leads to default values : int -> 0, string -> "", bool -> false
	var nums[4] int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	// nums[4] = 5 // This is a invalid argument -> index 4 out of bounds 

	fmt.Println(len(nums))
	fmt.Println(nums)

	var valss[4] bool
	valss[1] = true
	fmt.Println(valss)

	var names[3] string
	names[0] = "SaiDaiwik"
	names[1] = "Nanda"
	names[2] = "Yaswanth"
	fmt.Println(names)

	// to declare in single line
	num := [3]int{1,2,3}
	fmt.Println(num)

	// 2D-Arrays:
	n1 := [2][2]int{{1,2},{2,3}}
	fmt.Println(n1)

	// Array benefits :
		// - Fixed size, that is predictable
		// - Memory optimazation
		// - Constant time access
}