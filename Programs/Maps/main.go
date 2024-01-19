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
	if game.Name == "" {
		return game, errors.New("Name cannot be empty")
	}
	if game.Year <= 0 {
		return game, errors.New("Invalid year")
	}
	if game.Rating == "" {
		return game, errors.New("Rating cannot be empty")
	}
	if game.Price <= 0 {
		return game, errors.New("Invalid price")
	}

	return game, nil
}

func main() {
	game, err := collectGameInfo()

	if err != nil {
		fmt.Printf("Error collecting game information: %v\n", err)
		return
	}

	err = AddNewGame(game)
	if err != nil {
		fmt.Printf("Failed to add game: %v\n", err)
		return
	}

	fmt.Println("Game added:", game)
}
