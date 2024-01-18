package main

import (
	"errors"
	"fmt"
)

func ErrorGame() error {
	return errors.New("Error creating new game")
}

func userNewGame() (Games, error) {
	var game Games

	fmt.Println("Enter the name of the game: ")
	fmt.Scanln(&game.Name)

	fmt.Println("Enter the year the game was released: ")
	fmt.Scanln(&game.Year)

	fmt.Println("Enter the rating of the game: ")
	fmt.Scanln(&game.Rating)

	fmt.Println("Enter the price of the game: ")
	fmt.Scanln(&game.Price)

	// Add your validation logic here
	if game.Name == "" || game.Year <= 0 || game.Rating == "" || game.Price <= 0 {
		return game, ErrorGame()
	}

	return game, nil
}

func main() {
	game, err := userNewGame()

	if err != nil {
		fmt.Printf("%s: Enter valid inputs\n", ErrorGame())
		return
	}

	err = AddItem(game)
	if err != nil {
		fmt.Printf("%s: Failed to add game\n", ErrorGame())
		return
	}

	fmt.Println("Game added:", game)
}
