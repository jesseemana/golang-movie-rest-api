package main

import "fmt"

func main() {	
	my_number := 23
	var ptr = &my_number

	*ptr = *ptr * 8

	fmt.Println("Memory address for the pointer is:", ptr)
	fmt.Println("Pointer value is:", *ptr)
}
