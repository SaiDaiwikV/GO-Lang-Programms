package main

import (
	"fmt"
	"slices"
)

func main(){

	// slices called dynamic arrays 
	// most used construct in go-lang

	var nums[] int

	fmt.Println(nums)
	fmt.Println(nums == nil) // in CPP or JAVA it is called NULL but in GO-LANG it is called NIL
	fmt.Println(len(nums))

	var vals = make([]int, 2, 5) // 2 is the size of the array

	// capacity is maximum number of elements can fit
	fmt.Println(cap(vals))
	fmt.Println(vals)
	fmt.Println(vals == nil)

	vals = append(vals, 1)
	vals = append(vals, 2)
	vals = append(vals, 3)
	vals = append(vals, 4)
	fmt.Println(vals)
	fmt.Println(cap(vals))



	var arr = make([]int, 0, 5)

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)

	fmt.Println(arr)
	fmt.Println(cap(arr))
	fmt.Println(len(arr))

	// copy functions 

	var arr1 = make([]int, 0, 5)
	arr1 = append(arr1, 2)

	var arr2 = make([]int, len(arr1))

	fmt.Println(arr1, arr2)
	copy(arr2, arr1)

	fmt.Println(arr1 , arr2)

	// slice operator
	var arr3 = []int{1,2,3,4,5}

	fmt.Println(arr3[0:3])
	fmt.Println(arr3[:2])
	fmt.Println(arr3[1:4])

	// slice

	var num1 = []int{1, 2}
	var num2 = []int{1, 2}

	fmt.Println(slices.Equal(num1,num2))

	// 2D - Slices
	var num3 = [][]int{{1,2},{2,3}}
	fmt.Println(num3)

}