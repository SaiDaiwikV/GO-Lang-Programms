package main

import (
	"bufio"
	"os"
	"fmt"
)

func files2(){
	sourceFile, err := os.Open("example.txt")
	
	if err != nil{
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("newFile.txt")

	if err != nil{
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writter := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil{
			if err.Error() != "EOF"{
				panic(err)
			}
			break
		}
		e := writter.WriteByte(b)
		if e != nil{
			panic(e)
		}
	}

	writter.Flush()
	
	fmt.Println("Data Written Successfully")
}