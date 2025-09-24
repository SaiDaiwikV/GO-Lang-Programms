package main

import (
	"fmt"
)

func counter() func() int {
	count := 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()

	fmt.Println(increment()) // Call the function with ()
	fmt.Println(increment()) // Call it again to see the closure in action
	fmt.Println(increment()) // Each call increments the count
}
