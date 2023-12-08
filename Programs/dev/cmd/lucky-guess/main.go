package main

import (
	"fmt"
	"math/rand"

	"vs.local/pkg/util"
)

const (
	// the maximum number of turns that the game allows.
	maxTurns = 3
)

func main() {
	fmt.Printf("Let's play hot or cold, take guess from 0-9. You only got 3 tries\n")

	target := rand.Intn(10)
	playerHasWon := false

	for i := 0; i < maxTurns; i++ {
		var guess int

		fmt.Printf("Whats your guess? You have %d turns left\n", maxTurns-i)

		fmt.Scanln(&guess)

		if guess == target {
			playerHasWon = true
			break
		}

		d := util.Delta(guess, target)

		switch d {
		case 1:
			fmt.Println("You're hot as lava")
		case 2, 3, 4:
			fmt.Println("You're steaming hot")
		case 5, 6, 7, 8:
			fmt.Println("You're warm")
		case 9, 10:
			fmt.Println("You're cold")
		}
	}

	printWinStatus(playerHasWon)
}

// printWinStatus prints a message
func printWinStatus(hasWon bool) {
	if hasWon {
		fmt.Print("You won the challenge!")

		return
	}

	fmt.Print("You lost!")
}
