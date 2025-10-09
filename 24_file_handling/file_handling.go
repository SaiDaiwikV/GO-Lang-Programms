package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	f, err := os.Open("example.txt")

	if err != nil{
		// log the error
		panic(err)
	}

	fileInfo, err := f.Stat()
	
	if err != nil {
		// log the error
		panic(err)
	}

	fmt.Println("The file name is:",fileInfo.Name())
	fmt.Println("File or Folder:",fileInfo.IsDir())
	fmt.Println("File Size:",fileInfo.Size())
	fmt.Println("File Permissions:",fileInfo.Mode())
	fmt.Println("File Modified at:",fileInfo.ModTime())

	fmt.Println("----------------------------")

	f, err = os.Open("example.txt")

	if err != nil{
		panic(err)
	}
	defer f.Close()

	buff := make([]byte,12)

	d, err := f.Read(buff)

	if err != nil{
		panic(err)
	}

	for i:=0;i<len(buff);i++{
		fmt.Println("data:",d,string(buff[i]))
	}

	fileR, err := os.ReadFile("example.txt")

	if err != nil{
		panic(err)
	}
	

	fmt.Println(string(fileR))

	// read dir

	dir, err := os.Open(".")
	if err != nil{
		panic(err)
	}

	defer dir.Close()

	dirInfo,err := dir.ReadDir(0)

	for _, d := range dirInfo{
		fmt.Println(d.Name())
	}

	fmt.Println("-------------------")

	dir1, err := os.Open("../")
	if err != nil{
		panic(err)
	}

	defer dir1.Close()

	dirInfo, err = dir1.ReadDir(0)

	for _,d := range dirInfo{
		fmt.Println(d.Name()," - ",d.IsDir())
	}

	// create a file

	filec, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	defer filec.Close()

	filec.WriteString("Hi go")
	filec.WriteString(" nice language")

	dataEx2, err := os.ReadFile("example2.txt")

	if err != nil{
		panic(err)
	}

	content := string(dataEx2)
	fmt.Println("Before changing :",content)

	newContent := strings.Replace(content, "Hi go","Go-lang",1)

	err = os.WriteFile("example2.txt",[]byte(newContent), 0644)

	if err != nil{
		panic(err)
	}

	dataEx2, err = os.ReadFile("example2.txt")

	if err != nil {
		panic(err)
	}

	data := string(dataEx2)

	fmt.Println(data)

	fmt.Println("----End of file-handling.go--------below files1.go-code----")

	files1()

}
