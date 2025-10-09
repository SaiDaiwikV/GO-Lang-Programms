package main

import (
	"fmt"
	"os"
)

func file3(){
	// this code will delete a file:
	filename := "newFile.txt"

	err := os.Remove(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("File Deleted Successfully.")
}