package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	welcome := "Just some Golang stuff"
	fmt.Println(welcome)

	fmt.Println(multiply(87, 912))

	result := complex(87, 93, 32, 12, 93, 20)
	fmt.Println("Total input sum is", result)
}

func inputs() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your age?")
	input, _ :=  reader.ReadString('\n')

	fake_age, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You are now", fake_age + 1, "years old.")
}

func multiply(x int, y int) int {
	return x * y
}

func complex(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}  

	return total
}	
