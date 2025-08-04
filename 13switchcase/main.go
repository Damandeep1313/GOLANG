package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceNumber := rand.Intn(6) + 1
	fmt.Println("dice says : ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("you can open")
		fallthrough
	case 2:
		fmt.Println("you can move 2 steps")
		fallthrough
	case 3:
		fmt.Println("you can move 3 steps")
	
	case 4:
		fmt.Println("you can move 4 steps")
	
	case 5:
		fmt.Println("you can move 5 steps")
	
	case 6:
		fmt.Println("you can move 6 steps and roll again")
	
	default:
		fmt.Println("whats that eeew!")
	}	
	
}