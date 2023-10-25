package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Handling files in GO")

	content := "We've created a second file with Go. Here we are closing with defer."
	
	file, err := os.Create("./newfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("The length of the file is", length)	
	defer file.Close()
}	