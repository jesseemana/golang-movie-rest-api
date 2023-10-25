package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("The length of the file is", length)	
	defer file.Close()
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Data inside the file \n", string(data))
}	

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}	