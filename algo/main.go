package main

import "fmt"

func main() {
	s := ""
	for i := 1; i < 6; i++ {
		s += fmt.Sprint(algo(i))
	}

	fmt.Println(s)
}

func algo(n int) int {
	if n == 0 {
		return 2
	} else if n == 1 {
		return 1
	} else {
		return algo(n-1) + algo(n-2)
	}
}	