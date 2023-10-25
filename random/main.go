package main

import (
	"fmt"
	"time"
	"math/rand"
	// "math/big"
	// "crypto/rand"
)

func main() {
	fmt.Println("Random numbers first")
	// using Crypto
	// random_num, _ := rand.Int(rand.Reader, big.NewInt(7))
	// fmt.Println(random_num)

	current_time := time.Now()
	fmt.Println(current_time.Format("01-02-2006 15:04:05 Monday"))

	diceGame()
}

func diceGame() {
	rand.Seed(time.Now().UnixNano())
	dice_number := rand.Intn(6) + 1

	switch dice_number {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
		fallthrough
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What the heck bro")
	}
}
