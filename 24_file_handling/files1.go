package main

import (
	// "fmt"
	"os"
)

func files1(){
	filename := "example3.txt"

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes := []byte("Hello Golang")

	file.Write(bytes)

}
