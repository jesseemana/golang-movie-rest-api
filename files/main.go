package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Handling files in GO")

	writeFile()
	readFile("C:\\Users\\hey\\Desktop\\Code\\Golang\\Lessons\\files\\newfile.txt")
}	

func writeFile() {
	content := "We've created a second file with Go. Here we are closing with defer."
	
	file, err := os.Create("./newfile.txt")
	checkNilErr(err)

	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("The length of the file is", length)	
}

func readFile(filename string) {
	var response_string strings.Builder

	data, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	count, _ := response_string.Write(data)
	fmt.Println(count)

	fmt.Println("Data inside the file \n", response_string.String())
}	

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}	