package main

import "fmt"

func main() {
	s := ""
	for i := 1; i < 6; i++ {
		s += fmt.Sprint(algo(i))
	}

	fmt.Println(s)
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