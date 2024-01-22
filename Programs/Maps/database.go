package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

const (
	NewGame = "NewGame"
)

type Games struct {
	PK     string
	SK     string
	Name   string  // "Super Mario Bros", "Legend of Zelda", etc.
	Year   int     // 1998, 2001, etc.
	Rating string  // E, T, M, etc.
	Price  float64 // 59.99, 49.99, etc.
}

func newSess() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})
}

func generateRandomID() string {
	// Generate a new random UUID
	id := uuid.New()

	// Convert UUID to string
	idString := id.String()

	return idString
}

func getUserInput(prompt string) string {
	fmt.Print(prompt + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getIntInput(prompt string) (int, error) {
	inputStr := getUserInput(prompt)
	return strconv.Atoi(inputStr)
}

func getFloatInput(prompt string) (float64, error) {
	inputStr := getUserInput(prompt)

	// Remove leading "$" if present
	inputStr = strings.TrimLeft(inputStr, "$")

	// Attempt to parse the input as a float
	value, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		// If parsing as float fails, try parsing as an integer
		intValue, err := strconv.Atoi(inputStr)
		if err != nil {
			return 0, fmt.Errorf("invalid input: %v", err)
		}
		value = float64(intValue)
	}

	return value, nil
}

func collectGameInfo() (Games, error) {
	game := Games{
		Name: getUserInput("Enter the game name"),
	}

	// Get other fields from the user
	year, err := getIntInput("Enter the release year")
	if err != nil {
		fmt.Println("Invalid year")
		return game, err
	}
	game.Year = year

	rating := getUserInput("Enter the rating (E, T, M, etc.)")
	game.Rating = rating

	price, err := getFloatInput("Enter the price")
	if err != nil {
		fmt.Println("Invalid price")
		return game, err
	}

	game.Price = price

	return game, nil
}

func AddNewGame(Name string, Year int, Rating string, Price float64) error {
	sess, err := newSess()
	if err != nil {
		return err
	}

	db := dynamo.New(sess)
	table := db.Table("Maps-Example")

	randID := generateRandomID()

	PK := fmt.Sprintf("%s:NewGame", randID)

	newGame := Games{
		PK:     PK,
		SK:     NewGame,
		Name:   Name,
		Year:   Year,
		Rating: Rating,
		Price:  Price,
	}

	if err := table.Put(&newGame).Run(); err != nil {
		return fmt.Errorf("failed to add game: %v into the database", err)
	}

	return nil
}
