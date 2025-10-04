package main

import (
	"fmt"
	// "maps"
)

// iterating over data structures
func main() {
	nums := []int{4, 5, 6}
	sum := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}

	fmt.Println("The sum of the range is : ", sum)

	arr := []int{1, 2, 3, 4, 5}
	sum = 0

	for _, num := range arr {
		sum = sum + num
	}

	fmt.Println("The sum of arr is : ", sum)

	m := map[string]string{"Mobile": "POCO", "Age": "27", "Laptop": "Lenovo"}

	for k, v := range m {
		fmt.Println("The key is : ", k, " and the value is : ", v)
	}

	for i := range m{
		fmt.Println(i) // this will print only key of map -> "m"
	}


	for i, c := range "golang"{
		fmt.Println(i, c)
	}
	fmt.Println("-------------------------------")
	for i,c := range "goalng"{
		fmt.Println(i, string(c))
	}
}
