package main

import "fmt"

func main() {
	s := ""
	for i := 1; i < 6; i++ {
		s += algo(i)
	}

	fmt.Println(s)
	// defered things run at the end fo the function execution
	// defer fmt.Println("World")
	// fmt.Println("Hello")
}

func algo(value int) int {
	if value == 0 {
		return 2
	} else if value == 1 {
		return 1
	} else {
		return algo(value-1) + algo(value-2)
	}
}
